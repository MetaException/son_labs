package roundmanager

import (
	"lab1/internal/network"
	"lab1/internal/network/vertex"
	"strconv"
)

type RoundManager struct {
	G *network.Graph
}

func NewRoundManager(g *network.Graph) *RoundManager {
	return &RoundManager{
		G: g,
	}
}

// Выполняет раунды до тех пор, пока во всех узлах не закончатся кадры
func (r *RoundManager) PerformRounds() {
	i := 1
	for !r.CheckAllPoweroff() && !r.CheckFinished() {
		r.PerformRound(i)
		r.PerformMoving()

		r.G.VertexMap = make(map[*vertex.Vertex][]*vertex.Vertex)
		r.G.FillGraph()

		i++
	}
}

func (r *RoundManager) PerformRound(roundNumber int) {

	for _, sender := range r.G.Nodes {

		if sender.Power <= 0 {
			continue
		}

		recievers := r.G.VertexMap[&sender.Vertex]
		for i := range recievers {
			recievers[i] = sender.Vertex.Send(recievers[i], sender.FpR)
			if node, ok := r.G.Nodes[recievers[i].Name]; ok { // Получатель - node
				node.Power -= 0.2
			}
		}

		sender.DestroyFrames(sender.FpR)
		sender.Power--
		sender.R *= (sender.Power / 100)
	}

	r.G.PrintInfo(roundNumber)
	r.ClearAllDeadFramesHistory()

	r.G.DrawGraph(strconv.Itoa(roundNumber))
}

func (r *RoundManager) PerformMoving() {
	for _, node := range r.G.Nodes {
		node.RandomMove(float64(r.G.AreaX), float64(r.G.AreaY))
	}
}

func (r *RoundManager) ClearAllDeadFramesHistory() {
	for _, node := range r.G.VertexList {
		node.ClearDeadFramesHistory()
	}
}

func (r *RoundManager) CheckFinished() bool {

	for _, node := range r.G.Nodes {
		if len(node.Frames) != 0 {
			return false
		}
	}
	return true
}

func (r *RoundManager) CheckAllPoweroff() bool {

	for _, node := range r.G.Nodes {
		if node.Power > 0 {
			return false
		}
	}
	return true
}
