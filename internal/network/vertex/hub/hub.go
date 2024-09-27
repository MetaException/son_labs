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

func GenerateRandomHubByBaseNode(name string, src vertex.Vertex) *Hub {

	base := vertex.GenerateRandomBaseNode(name, src)
	base.X = 50
	base.Y = 50

	fmt.Printf("New hub [%s] : X: %v, Y: %v, R: %v\n", name, base.X, base.Y, base.R)

	return &Hub{
		Vertex: *base,
	}
}
