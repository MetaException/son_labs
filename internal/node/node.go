package node

import (
	"lab1/internal/frame"
	"lab1/pkg/utils"
	"math"
	"strconv"
)

// Узел в сети
type Node struct {
	BaseNode
	FpR             int
	FramesIdHistory map[string]int
	MovingSpeed     float64
}

func (node Node) String() string {
	return node.BaseNode.Name
}

func NewNode(X, Y, R float64, FpR int, Name string, frameCount int) *Node {
	node := &Node{
		BaseNode:        *NewBaseNode(X, Y, R, Name),
		FpR:             FpR,
		FramesIdHistory: make(map[string]int),
		MovingSpeed:     10,
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

	base := GenerateRandomBaseNode(name, s.BaseNode)
	nodeFrameCount := utils.GenerateRandomInt(1, 3)
	fpr := utils.GenerateRandomInt(0, 5)

	//fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, int(x), int(y), rr, nodeFrameCount)

	return NewNode(base.X, base.Y, base.R, fpr, name, nodeFrameCount)
}

func (s Node) IsAdjacent(vertexToCompare *Node) bool {
	return math.Sqrt(math.Pow(vertexToCompare.BaseNode.X-s.BaseNode.X, 2)+math.Pow(vertexToCompare.BaseNode.Y-s.BaseNode.Y, 2)) <= math.Max(s.BaseNode.R, vertexToCompare.BaseNode.R)
}
