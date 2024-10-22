package hub

import (
	"lab1/internal/network/vertex"
)

type Hub struct {
	vertex.Vertex
}

func (n *Hub) UpdateBase(base *vertex.Vertex) {
	n.Vertex = *base
}
