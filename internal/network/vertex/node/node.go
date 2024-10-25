package node

import (
	"lab1/internal/network/frame"
	"lab1/internal/network/vertex"
	"strconv"
)

// Узел в сети
type Node struct {
	vertex.Vertex
	FpR             int
	FramesIdHistory map[string]int
	MovingSpeed     float64
	Power           float64 //percent
}

func NewNode(vertex *vertex.Vertex, FpR int, frameCount int) *Node {
	node := &Node{
		Vertex:          *vertex,
		FpR:             FpR,
		FramesIdHistory: make(map[string]int),
		MovingSpeed:     3,
		Power:           100,
	}

	for i := range frameCount {
		frame := &frame.Frame{
			ParentName: vertex.Name,
			TTL:        100, // Ставить динамически
			ID:         node.Name + "-" + strconv.Itoa(i),
		}

		node.Frames = append(node.Frames, frame)
		node.FramesIdHistory[frame.ID] = frame.TTL
	}

	return node
}

func (node Node) String() string {
	return node.Vertex.Name
}

func (n *Node) UpdateBase(base *vertex.Vertex) {
	n.Vertex = *base
}
