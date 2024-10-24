package hub

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/pkg/utils"
)

func GenerateRandomHub(name string) *Hub {
	base := vertex.GenerateRandomBase(name)
	base.R = utils.GenerateRandom(5, 20)

	fmt.Printf("New hub [%s] : X: %v, Y: %v, R: %v\n", name, base.X, base.Y, base.R)

	return &Hub{
		Vertex: *base,
	}
}

func GenerateRandomHubByVertex(name string, source vertex.Vertex) *Hub {
	base := vertex.GenerateRandomBaseByVertex(name, source)
	base.R = utils.GenerateRandom(5, 20)

	fmt.Printf("New HUB [%s] : X: %v, Y: %v, R: %v\n", name, base.X, base.Y, base.R)

	return &Hub{
		Vertex: *base,
	}
}
