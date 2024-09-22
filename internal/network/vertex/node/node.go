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
	Power           byte //percent
}

func (node Node) String() string {
	return node.Vertex.Name
}

func NewNode(X, Y, R float64, FpR int, Name string, frameCount int) *Node {
	node := &Node{
		Vertex:          *vertex.NewBaseNode(X, Y, R, Name),
		FpR:             FpR,
		FramesIdHistory: make(map[string]int),
		MovingSpeed:     10,
		Power:           100,
	}

	for i := range frameCount {
		frame := &frame.Frame{
			ParentName: Name,
			TTL:        100, // Ставить динамически
			ID:         node.Name + "-" + strconv.Itoa(i),
		}

		node.Frames = append(node.Frames, frame)
		node.FramesIdHistory[frame.ID] = frame.TTL
	}

	return node
}

func (s *Node) GenerateRandomVertexByVertex(name string) *Node {

	base := vertex.GenerateRandomBaseNode(name, s.Vertex)
	nodeFrameCount := utils.GenerateRandomInt(1, 3)
	fpr := utils.GenerateRandomInt(1, 5)

	fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, base.X, base.Y, base.R, nodeFrameCount)

	return NewNode(base.X, base.Y, base.R, fpr, name, nodeFrameCount)
}
