package network

import (
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/node"
	"math"
)

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

	//	fmt.Printf("\n%s %v sends %v to %s %v", src.Name, src.Frames, framesToSend, reciever.Name, reciever.Frames)

	if node, ok := dist.(*node.Node); ok {
		node.Power = math.Max(0, node.Power-0.01*float64(sentCount))
		node.R *= (node.Power / 100)
	}
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
