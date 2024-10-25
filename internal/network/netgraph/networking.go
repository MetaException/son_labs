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

type hel struct {
	src vertex.IVertex
	n   vertex.IVertex
}

func (graph *Graph) PerformRounds(roundNumber int) {

	graph.ClearMap()
	graph.ClearRouteMap()
	graph.Fill(roundNumber)

	for _, nod := range graph.Nodes {

		neighbors := graph.VertexMap[nod] // Его соседи

		if nod.Power <= 0 {
			continue
		}

		bestAntPath := make(map[*node.Node][]hel, 0) // Лучшие пути муравьёв
		bestAntCost := math.MaxInt                   // Лучшие затраты на путь

		// Запускаем несколько агентов (муравьев)
		antsCount := len(graph.VertexMap[nod]) * 2 // Ставим количество муравьёв

		for antIndex := 0; antIndex < antsCount; antIndex++ {

			agentHistory := make([]hel, 0)               // История пути муравья
			agentHistoryMap := make(map[string]struct{}) // История посещённых узлов для исключения повторений
			agentHistoryMap[nod.Name] = struct{}{}

			fmt.Println("\n------")
			fmt.Printf("Ход муравья... Узел: %v | Соседи: %v\n", nod.Name, neighbors)

			currNode := nod            // Текущий узел муравья
			currNeighbors := neighbors // Текущие соседи узла
			for {                      // Пока муравей не дойдёт до выхода

				if _, ok := graph.RouteMap[currNode]; !ok {
					graph.RouteMap[currNode] = make(map[vertex.IVertex]*RoutingData)
				}

				choosedOne, err := graph.chooseNeighbor(currNode, currNeighbors, agentHistoryMap)
				if err != nil { // Не нашлось кандидатов - идём на один узел назад
					if len(agentHistory) == 0 {
						break
					}
					currNode = agentHistory[len(agentHistory)-1].src.(*node.Node)
					agentHistory = agentHistory[:len(agentHistory)-1]
					continue
				}

				agentHistory = append(agentHistory, hel{
					src: currNode,
					n:   choosedOne,
				}) // Добавляем в историю выбранный узел
				agentHistoryMap[choosedOne.GetBase().Name] = struct{}{} // Добавляем в историю посещённых узлов

				fmt.Printf("Выбран сосед: %v | История соседей: %v\n", choosedOne, agentHistory)

				if _, ok := choosedOne.(*hub.Hub); ok {
					break
				}

				currNode = choosedOne.(*node.Node)        // Муравей перешёл на следующий узел
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

		graph.clPh()

		fmt.Println("Обновляем феромоны...")
		// Добавляем феромоны на пути, по которому прошёл лучший муравей
		for _, v := range bestAntPath {
			for i := len(v) - 1; i >= 0; i-- {
				rd := graph.RouteMap[v[i].src][v[i].n]
				rd.Pintensity += (1-0.2)*rd.Pintensity + (1.0 / float64(bestAntCost))
				rd.Cost = float64(bestAntCost)
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

func (graph *Graph) clPh() {
	for _, v := range graph.RouteList {
		v.Pintensity *= (1 - 0.05)
	}
}

func (graph *Graph) chooseNeighbor(currNode *node.Node, potentialNeighbors []vertex.IVertex, agentHistoryMap map[string]struct{}) (vertex.IVertex, error) {

	totalPheromones := 0.0 // Общее число феромонов
	alpha := 1.0           // Коэффициент усиления влияния феромонов

	neighbors := make([]*node.Node, 0)
	for _, neighbor := range potentialNeighbors {
		if _, ok := agentHistoryMap[neighbor.GetBase().Name]; !ok {

			var RD *RoutingData

			// Формируем карту феромонов
			if rd, isExist := graph.RouteMap[currNode][neighbor]; !isExist {
				RD = &RoutingData{
					Cost:       0.0,
					Pintensity: 0.0,
				}
				graph.RouteMap[currNode][neighbor] = RD
				if _, ok := graph.RouteMap[neighbor]; !ok {
					graph.RouteMap[neighbor] = make(map[vertex.IVertex]*RoutingData)
				}
				graph.RouteMap[neighbor][currNode] = RD
				graph.RouteList = append(graph.RouteList, RD)
			} else {
				RD = rd
			}

			if hub, isHub := graph.Hubs[neighbor.GetBase().Name]; isHub {
				return hub, nil
			} else if node, isNode := graph.Nodes[neighbor.GetBase().Name]; isNode { // Значит это нод
				totalPheromones += math.Pow(RD.Pintensity, alpha) / (RD.Cost + 1) // Рассчитываем сумму феромонов для всех соседей
				neighbors = append(neighbors, node)
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

			if neighbor.Power <= 0 {
				continue
			}

			rd := graph.RouteMap[currNode][neighbor]
			probability := (math.Pow(rd.Pintensity, alpha) / (rd.Cost + 1)) / totalPheromones
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
