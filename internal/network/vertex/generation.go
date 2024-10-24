package vertex

import (
	"lab1/pkg/utils"
	"math"
)

func GenerateRandomBase(name string) *Vertex {
	x := float64(utils.GenerateRandomInt(0, 100))
	y := float64(utils.GenerateRandomInt(0, 100))
	r := float64(utils.GenerateRandomInt(5, 30))

	return NewBaseNode(float64(int(x)), float64(int(y)), r, name)
}

func GenerateRandomBaseByVertex(name string, source Vertex) *Vertex {

	base := GenerateRandomBase(name)

	cx := source.X
	cy := source.Y

	dx := base.X - cx
	dy := base.Y - cy
	distance := math.Sqrt(dx*dx + dy*dy)

	cr := base.R

	if distance > cr {
		ratio := cr / distance
		base.X = cx + dx*ratio
		base.Y = cy + dy*ratio
	}

	return base
}
