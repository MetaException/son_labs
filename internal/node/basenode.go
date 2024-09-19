package node

import (
	"lab1/internal/frame"
	"lab1/pkg/utils"
	"math"
)

type BaseNode struct {
	X, Y, R         float64
	Name            string
	Frames          []*frame.Frame // Сделать лимит??
	FramesIdHistory map[string]int
}

func NewBaseNode(X, Y, R float64, Name string) *BaseNode {
	return &BaseNode{
		X:               X,
		Y:               Y,
		R:               R,
		Name:            Name,
		Frames:          make([]*frame.Frame, 0),
		FramesIdHistory: make(map[string]int),
	}
}

func GenerateRandomBaseNode(name string, base BaseNode) *BaseNode {

	x := float64(utils.GenerateRandomInt(0, 100))
	y := float64(utils.GenerateRandomInt(0, 100))

	cx := base.X
	cy := base.Y

	dx := x - cx
	dy := y - cy
	distance := math.Sqrt(dx*dx + dy*dy)

	cr := base.R

	if distance > cr {
		ratio := cr / distance
		x = cx + dx*ratio
		y = cy + dy*ratio
	}

	r := utils.GenerateRandom(5, 50)

	//fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, int(x), int(y), rr, nodeFrameCount)

	return NewBaseNode(x, y, r, name)

	//return NewNode(x, y, rr, 2, name, nodeFrameCount)
}
