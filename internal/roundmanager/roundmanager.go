package roundmanager

import (
	"lab1/internal/network"
	"lab1/internal/node"
)

type RoundManager struct {
	G *network.Graph
}

func NewRoundManager(g *network.Graph) *RoundManager {
	return &RoundManager{
		G: g,
	}
}

// Выполняет N раундов
func (r *RoundManager) PerformNRounds(count int) {
	for i := range count {
		r.PerformRound(i + 1)
		r.PerformMoving()
	}
}

// Выполняет раунды до тех пор, пока hub не достингент определённого количества кадров
func (r *RoundManager) PerformRoundUntilFrameCount(count int) {
	i := 1
	for len(r.G.VertexByName["hub"].Frames) != count {
		r.PerformRound(i)
		r.PerformMoving()
		i++
	}
}

// Выполняет раунды до тех пор, пока во всех узлах не закончатся кадры
func (r *RoundManager) PerformRounds() {
	i := 1
	for !r.CheckFinished() {
		r.PerformRound(i)
		r.PerformMoving()

		r.G.VertexMap = make(map[*node.Node][]*node.Node)
		r.G.FillGraph()

		i++
	}
}

func (r *RoundManager) PerformRound(roundNumber int) {

	for _, sender := range r.G.VertexList {

		if sender.IsHub {
			continue
		}

		recievers := r.G.VertexMap[sender]
		for i := range recievers {
			recievers[i] = sender.Send(recievers[i], sender.FpR)
		}

		sender.DestroyFrames(sender.FpR)
	}

	r.G.PrintInfo(roundNumber)
	r.ClearAllDeadFramesHistory()
}

func (r *RoundManager) PerformMoving() {
	for _, node := range r.G.VertexList {
		if node.IsHub {
			continue
		}

		node.RandomMove(float64(r.G.AreaX), float64(r.G.AreaY))
	}
}

func (r *RoundManager) ClearAllDeadFramesHistory() {
	for _, node := range r.G.VertexList {
		node.ClearDeadFramesHistory()
	}
}

func (r *RoundManager) CheckFinished() bool {

	for _, node := range r.G.VertexList {
		if node.IsHub {
			continue
		}
		if len(node.Frames) != 0 {
			return false
		}
	}

	return true
}
