package vertex

import (
	"lab1/pkg/utils"
	"math"
)

func GenerateRandomBaseNode(name string, base Vertex) *Vertex {

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

	r := float64(utils.GenerateRandomInt(5, 50))

	return NewBaseNode(x, y, r, name)
}

func (s Vertex) IsAdjacent(vertexToCompare *Vertex) bool {
	return math.Sqrt(math.Pow(vertexToCompare.X-s.X, 2)+math.Pow(vertexToCompare.Y-s.Y, 2)) <= math.Max(s.R, vertexToCompare.R)
}
