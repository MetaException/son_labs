package hub

import "lab1/internal/node"

type Hub struct {
	node.BaseNode
}

func NewHub() *Hub {
	return &Hub{}
}

func GenerateRandomHubByBaseNode(name string, vertex node.BaseNode) *Hub {

	base := node.GenerateRandomBaseNode(name, vertex)

	return &Hub{
		BaseNode: *base,
	}
}
