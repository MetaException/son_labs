package network

import (
	"fmt"
	"lab1/internal/hub"
	"lab1/internal/node"
)

type Graph struct {
	VertexMap    map[*node.Node][]*node.Node // Список смежности
	VertexList   []*node.Node
	VertexByName map[string]*node.Node

	Hubs []*hub.Hub

	Length       int
	AreaX, AreaY int
}

func NewGraph(N int) *Graph {
	return &Graph{
		VertexMap:    make(map[*node.Node][]*node.Node, N),
		VertexList:   make([]*node.Node, 0),
		Length:       N,
		VertexByName: make(map[string]*node.Node),
		AreaX:        100,
		AreaY:        100,
		Hubs:         make([]*hub.Hub, 0),
	}
}

func (g *Graph) AddNode(nodeToAdd *node.Node) {

	if _, ok := g.VertexMap[nodeToAdd]; !ok {
		g.VertexMap[nodeToAdd] = make([]*node.Node, 0)
		g.VertexList = append(g.VertexList, nodeToAdd)
		g.VertexByName[nodeToAdd.Name] = nodeToAdd
	}
}

func (g *Graph) AddEdge(vertex *node.Node, adjacentVertex *node.Node) {

	g.VertexMap[vertex] = append(g.VertexMap[vertex], adjacentVertex)
	g.VertexMap[adjacentVertex] = append(g.VertexMap[adjacentVertex], vertex)
}

func (g *Graph) FillGraph() {

	for i := 0; i < len(g.VertexList); i++ {
		for j := i + 1; j < len(g.VertexList); j++ {

			vertex := g.VertexList[i]
			vertexToCompare := g.VertexList[j]

			if vertex.IsAdjacent(vertexToCompare) {
				g.AddEdge(vertex, vertexToCompare)
			}
		}
	}

	fmt.Println(g.VertexMap)
}
