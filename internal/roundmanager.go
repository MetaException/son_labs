package roundmanager

import (
	"lab1/internal/network"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/node"
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

		r.G.DrawGraph(strconv.Itoa(i))

		r.G.VertexMap = make(map[vertex.IVertex][]vertex.IVertex)
		r.G.FillGraph()

		i++
	}
}

func (r *RoundManager) PerformRound(roundNumber int) {

	for _, sender := range r.G.Nodes {

		if sender.Power <= 0 {
			continue
		}

		recievers := r.G.VertexMap[sender]
		for i := range recievers {
			Communicate(sender, recievers[i], sender.FpR)
		}

		sender.DestroyFrames(sender.FpR)
		sender.Power--
		sender.R *= (sender.Power / 100)
	}

	r.G.PrintInfo(roundNumber)
	r.ClearAllDeadFramesHistory()
}

func (r *RoundManager) PerformMoving() {
	for _, node := range r.G.Nodes {
		node.RandomMove(float64(r.G.AreaX), float64(r.G.AreaY))
	}
}

func (r *RoundManager) ClearAllDeadFramesHistory() {
	for _, vertex := range r.G.VertexList {
		vertex.ClearDeadFramesHistory()
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

func Communicate(src *node.Node, dist vertex.IVertex, count int) { // TODO: убрать в другое место??

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

	//	fmt.Printf("\n%s %v sends %v to %s %v", src.Name, src.Frames, framesToSend, reciever.Name, reciever.Frames)

	if node, ok := dist.(*node.Node); ok {
		node.Power -= 0.2 * float64(sentCount)
	}
}
