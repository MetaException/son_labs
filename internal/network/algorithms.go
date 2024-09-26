package network

import (
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/node"
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
		node.Power -= 0.2 * float64(sentCount)
	}
}
