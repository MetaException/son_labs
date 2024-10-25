package netgraph

import (
	"errors"
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"math"
	"math/rand"
)

func (graph *Graph) PerformRounds(roundNumber int) {

	graph.ClearMap()
	graph.resetPh()
	graph.Fill(roundNumber)

	for _, nod := range graph.Nodes {

		neighbors := graph.VertexMap[nod] // Его соседи

		if len(graph.VertexMap[graph.Hubs["hub"]]) == 0 {
			break
		}

		if len(nod.Frames) == 0 || nod.Power <= 0 {
			continue
		}

		bestAntPath := make(map[*node.Node][]*pNeighbor, 0) // Лучшие пути муравьёв
		bestAntCost := math.MaxInt                          // Лучшие затраты на путь

		// Запускаем несколько агентов (муравьев)
		antsCount := len(graph.VertexMap[nod]) * 2 // Ставим количество муравьёв

		for antIndex := 0; antIndex < antsCount; antIndex++ {

			agentHistory := make([]*pNeighbor, 0)   // История пути муравья
			agentHistoryMap := map[string]struct{}{ // История посещённых узлов для исключения повторений
				nod.Name: {},
			}

			fmt.Println("\n------")
			fmt.Printf("Ход муравья... Узел: %v | Соседи: %v\n", nod.Name, neighbors)

			currNode := nod            // Текущий узел муравья
			currNeighbors := neighbors // Текущие соседи узла
			for {                      // Пока муравей не дойдёт до выхода
				choosedOne, err := graph.chooseNeighbor(currNode, currNeighbors, agentHistoryMap)
				if err != nil { // Не нашлось кандидатов - идём на один узел назад
					if len(agentHistory) == 0 {
						break
					}
					currNode = agentHistory[len(agentHistory)-1].src
					agentHistory = agentHistory[:len(agentHistory)-1]
					continue
				}

				agentHistory = append(agentHistory, choosedOne)              // Добавляем в историю выбранный узел
				agentHistoryMap[choosedOne.node.GetBase().Name] = struct{}{} // Добавляем в историю посещённых узлов

				fmt.Printf("Выбран сосед: %v | История соседей: %v\n", choosedOne, agentHistory)

				if _, ok := choosedOne.node.(*hub.Hub); ok {
					break
				}

				currNode = choosedOne.node.(*node.Node)   // Муравей перешёл на следующий узел
				currNeighbors = graph.VertexMap[currNode] // Получаем соседей уже этого узла
			}

			fmt.Println("Достигли хаба")
			// Оценка текущего пути муравья
			currentAntCost := len(agentHistory) // Длина пути как стоимость
			if currentAntCost > 0 && currentAntCost < bestAntCost {
				bestAntCost = currentAntCost
				for k := range bestAntPath {
					delete(bestAntPath, k)
				}
				bestAntPath[currNode] = agentHistory
			}
		}

		// Испарение феромонов
		graph.clPh()

		fmt.Println("Обновляем феромоны...")
		// Добавляем феромоны на пути, по которому прошёл лучший муравей
		for _, v := range bestAntPath {
			for i := len(v) - 1; i >= 0; i-- {
				/* // Если что-то не будет работать!!!!
				if v[i].node.GetBase().Name != graph.VertexMap[v[i].src][v[i].idx].GetBase().Name {
					panic("rgj")
				}
				*/
				rd := v[i].node
				base := rd.GetBase()
				base.Pintensity += (1-0.2)*base.Pintensity + (1.0 / float64(bestAntCost))
				base.Cost = float64(bestAntCost)
			}
		}

		fmt.Printf("Лучшие пути: %v", bestAntPath)

		// Обрабатываем кадры
		Flooding(nod, graph.Hubs["hub"], nod.FpR)
		nod.DestroyFrames(nod.FpR)
		nod.Power--
		if nod.Power < 0 {
			nod.Power = 0
		}
	}

	graph.PrintInfo(roundNumber)
	graph.ClearAllDeadFramesHistory()
}

func (graph *Graph) resetPh() {
	for _, v := range graph.VertexList {
		v.GetBase().Pintensity = 0.0
	}
}

func (graph *Graph) clPh() {
	for _, v := range graph.VertexList {
		v.GetBase().Pintensity *= (1 - 0.45)
	}
}

type pNeighbor struct {
	node vertex.IVertex
	src  *node.Node
	idx  int
}

// !!!! Пропускаются соседи с зарядом 0

func (graph *Graph) chooseNeighbor(currNode *node.Node, potentialNeighbors []vertex.IVertex, agentHistoryMap map[string]struct{}) (*pNeighbor, error) {

	totalPheromones := 0.0 // Общее число феромонов
	alpha := 1.1           // Коэффициент усиления влияния феромонов

	neighbors := make([]*pNeighbor, 0)
	for i, neighbor := range potentialNeighbors {
		base := neighbor.GetBase()
		if _, ok := agentHistoryMap[base.Name]; !ok {
			if hub, isHub := graph.Hubs[base.Name]; isHub {
				return &pNeighbor{
					src:  currNode,
					node: hub,
					idx:  i,
				}, nil
			} else if node, isNode := graph.Nodes[base.Name]; isNode && node.Power > 0 {
				totalPheromones += math.Pow(base.Pintensity, alpha) / (base.Cost + 1) // Рассчитываем сумму феромонов для всех соседей
				neighbors = append(neighbors, &pNeighbor{
					src:  currNode,
					node: node,
					idx:  i,
				})
			}
		}
	}

	if len(neighbors) == 0 {
		return nil, errors.New("нет доступных кандидатов")
	}

	// Порог для активации детерминированного выбора
	threshold := 0.1
	explorationProbability := 0.1 // Вероятность для исследования

	// Если феромоны достаточны, выбираем лучший путь
	if totalPheromones > threshold {
		randValue := rand.Float64()
		cumulativeProbability := 0.0

		for _, neighbor := range neighbors {

			rd := graph.VertexMap[currNode][neighbor.idx]
			base := rd.GetBase()
			probability := (math.Pow(base.Pintensity, alpha) / (base.Cost + 1)) / totalPheromones
			cumulativeProbability += probability

			if randValue < cumulativeProbability {
				return neighbor, nil
			}
		}
	}

	// Если феромонов недостаточно или есть редкая необходимость исследования
	if rand.Float64() < explorationProbability {
		fmt.Println("Исследуем новый путь")
		return neighbors[rand.Intn(len(neighbors))], nil
	}

	// По умолчанию возвращаем первый сосед, если феромоны равны 0 или исследование неактивно
	return neighbors[0], nil
}

func (graph *Graph) PerformMoving() {
	for _, node := range graph.Nodes {
		node.RandomMove(float64(graph.AreaX), float64(graph.AreaY))
	}
}

func Flooding(src *node.Node, dist vertex.IVertex, count int) { // TODO: убрать в другое место??

	if count > len(src.Vertex.Frames) {
		count = len(src.Vertex.Frames)
	}

	framesToSend := src.Vertex.Frames[:count]

	sentCount := 0.0
	reciever := dist.GetBase()

	for _, frame := range framesToSend {
		if frame.ParentName != reciever.Name {
			if _, ok := reciever.FramesIdHistory[frame.ID]; !ok && frame.TTL > 0 {
				reciever.FramesIdHistory[frame.ID] = frame.TTL
				reciever.Frames = append(reciever.Frames, frame)
				frame.TTL--
				sentCount++
			}
		}
	}

	src.Power = math.Max(0, src.Power-0.01*float64(sentCount))
	src.R *= 1 - (src.Power/100)/500
}

func (graph *Graph) ClearAllDeadFramesHistory() {
	for _, vertex := range graph.VertexList {
		vertex.ClearDeadFramesHistory()
	}
}

func (graph *Graph) CheckFinished() bool {

	for _, node := range graph.Nodes {
		if len(node.Frames) != 0 {
			return false
		}
	}
	return true
}

func (graph *Graph) CheckAllPoweroff() bool {

	for _, node := range graph.Nodes {
		if node.Power > 0 {
			return false
		}
	}
	return true
}
