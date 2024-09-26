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

	//r := float64(utils.GenerateRandomInt(5, 50))

	return NewBaseNode(float64(int(x)), float64(int(y)), 100.0, name)
}

func IsAdjacent(ivertexSrc IVertex, ivertexToCompare IVertex) bool {

	vertexSrc := ivertexSrc.GetBase()
	vertexToCompare := ivertexToCompare.GetBase()

	return math.Sqrt(math.Pow(vertexToCompare.X-vertexSrc.X, 2)+math.Pow(vertexToCompare.Y-vertexSrc.Y, 2)) <= math.Max(vertexSrc.R, vertexToCompare.R)
}
