package node

import (
	"lab1/pkg/utils"
	"math"
)

func (s *Node) RandomMove(AreaX, AreaY float64) {

	leftBorderX := math.Max(s.Vertex.X-s.MovingSpeed, 0)
	rightBorderX := math.Min(s.Vertex.X+s.MovingSpeed, AreaX)

	leftBorderY := math.Max(s.Vertex.Y-s.MovingSpeed, 0)
	rightBorderY := math.Min(s.Vertex.Y+s.MovingSpeed, AreaY)

	newX := utils.GenerateRandom(leftBorderX, rightBorderX)
	newY := utils.GenerateRandom(leftBorderY, rightBorderY)

	s.Vertex.X = newX
	s.Vertex.Y = newY

	//fmt.Printf("Node %s (%v;%v) | Goto: x: %v, y: %v\n", s.Name, s.X, s.Y, newX, newY)
}
