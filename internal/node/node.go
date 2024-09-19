package node

import (
	"fmt"
	"lab1/internal/frame"
	"lab1/pkg/utils"
	"math"
	"strconv"
)

// Узел в сети
type Node struct {
	x, y, r         float64
	FpR             int
	Name            string
	Frames          []*frame.Frame // Сделать лимит??
	IsHub           bool
	FramesIdHistory map[string]int
	MovingSpeed     float64
}

func (node Node) String() string {
	return node.Name
}

func NewNode(X, Y, R, FpR int, Name string, isHub bool, frameCount int) *Node {
	node := &Node{
		x:               float64(X),
		y:               float64(Y),
		r:               float64(R),
		FpR:             FpR,
		Name:            Name,
		IsHub:           isHub,
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

func (s *Node) GenerateRandomVertexByVertex(name string, isHub bool) *Node {

	x := float64(utils.GenerateRandomInt(0, 100))
	y := float64(utils.GenerateRandomInt(0, 100))

	cx := s.x
	cy := s.y

	dx := x - cx
	dy := y - cy
	distance := math.Sqrt(dx*dx + dy*dy)

	r := s.r

	if distance > r {
		ratio := r / distance
		x = cx + dx*ratio
		y = cy + dy*ratio
	}

	// Генерируем случайный радиус и количество кадров
	rr := utils.GenerateRandomInt(5, 50)
	nodeFrameCount := utils.GenerateRandomInt(1, 3)

	if isHub {
		nodeFrameCount = 0
	}

	fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, int(x), int(y), rr, nodeFrameCount)

	return NewNode(int(x), int(y), rr, 2, name, isHub, nodeFrameCount)
}

func (s Node) IsAdjacent(vertexToCompare *Node) bool {
	return math.Sqrt(math.Pow(vertexToCompare.x-s.x, 2)+math.Pow(vertexToCompare.y-s.y, 2)) <= math.Max(s.r, vertexToCompare.r)
}
