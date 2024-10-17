package node

import (
	"fmt"
	"lab1/internal/network/frame"
	"lab1/internal/network/vertex"
	"lab1/pkg/utils"
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

func (node Node) String() string {
	return node.Vertex.Name
}

func (n *Node) UpdateBase(base *vertex.Vertex) {
	n.Vertex = *base
}

func GenerateRandomNode(name string) *Node {
	base := vertex.GenerateRandomBase(name)
	nodeFrameCount := utils.GenerateRandomInt(5, 10)
	fpr := utils.GenerateRandomInt(1, 5)

	fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, base.X, base.Y, base.R, nodeFrameCount)

	return NewNode(base, fpr, nodeFrameCount)
}

func GenerateRandomNodeByVertex(name string, source vertex.Vertex) *Node {
	base := vertex.GenerateRandomBaseByVertex(name, source)
	nodeFrameCount := utils.GenerateRandomInt(5, 10)
	fpr := utils.GenerateRandomInt(1, 5)

	fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, base.X, base.Y, base.R, nodeFrameCount)

	return NewNode(base, fpr, nodeFrameCount)
}

func NewNode(vertex *vertex.Vertex, FpR int, frameCount int) *Node {
	node := &Node{
		Vertex:          *vertex,
		FpR:             FpR,
		FramesIdHistory: make(map[string]int),
		MovingSpeed:     5,
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
