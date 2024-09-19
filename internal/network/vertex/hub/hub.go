package hub

import (
	"fmt"
	"lab1/internal/network/vertex"
)

type Hub struct {
	vertex.Vertex
}

func NewHub() *Hub {
	return &Hub{}
}

func GenerateRandomHubByBaseNode(name string, src vertex.Vertex) *Hub {

	base := vertex.GenerateRandomBaseNode(name, src)

	fmt.Printf("New hub [%s] : X: %v, Y: %v, R: %v\n", name, base.X, base.Y, base.R)

	return &Hub{
		Vertex: *base,
	}
}
