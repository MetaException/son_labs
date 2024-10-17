package hub

import (
	"fmt"
	"lab1/internal/network/vertex"
)

type Hub struct {
	vertex.Vertex
}

func (n *Hub) UpdateBase(base *vertex.Vertex) {
	n.Vertex = *base
}

func GenerateRandomHub(name string) *Hub {
	base := vertex.GenerateRandomBase(name)

	fmt.Printf("New hub [%s] : X: %v, Y: %v, R: %v\n", name, base.X, base.Y, base.R)

	return &Hub{
		Vertex: *base,
	}
}

func GenerateRandomHubByVertex(name string, source vertex.Vertex) *Hub {
	base := vertex.GenerateRandomBaseByVertex(name, source)

	fmt.Printf("New HUB [%s] : X: %v, Y: %v, R: %v\n", name, base.X, base.Y, base.R)

	return &Hub{
		Vertex: *base,
	}
}
