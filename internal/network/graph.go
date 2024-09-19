package network

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
)

type Graph struct {
	VertexMap    map[*vertex.Vertex][]*vertex.Vertex // Список смежности
	VertexByName map[string]*vertex.Vertex           // Доступ к вершинам по имени
	VertexList   []*vertex.Vertex                    // Все вершины графа сети

	Nodes map[string]*node.Node
	Hubs  map[string]*hub.Hub

	Length       int
	AreaX, AreaY int
}

func NewGraph(N int) *Graph {
	return &Graph{
		VertexMap:    make(map[*vertex.Vertex][]*vertex.Vertex, N),
		Nodes:        make(map[string]*node.Node, 0),
		Length:       N,
		VertexByName: make(map[string]*vertex.Vertex),
		AreaX:        100,
		AreaY:        100,
		Hubs:         make(map[string]*hub.Hub),
		VertexList:   make([]*vertex.Vertex, 0),
	}
}

func (g *Graph) AddHub(hubToAdd *hub.Hub) {
	if _, ok := g.VertexMap[&hubToAdd.Vertex]; !ok {
		g.VertexMap[&hubToAdd.Vertex] = make([]*vertex.Vertex, 0)
		g.Hubs[hubToAdd.Vertex.Name] = hubToAdd
		g.VertexList = append(g.VertexList, &hubToAdd.Vertex)
	}
}

func (g *Graph) AddNode(nodeToAdd *node.Node) {

	if _, ok := g.VertexMap[&nodeToAdd.Vertex]; !ok {
		g.VertexMap[&nodeToAdd.Vertex] = make([]*vertex.Vertex, 0)
		g.Nodes[nodeToAdd.Name] = nodeToAdd
		g.VertexByName[nodeToAdd.Name] = &nodeToAdd.Vertex
		g.VertexList = append(g.VertexList, &nodeToAdd.Vertex)
	}
}

func (g *Graph) AddEdge(vertex *vertex.Vertex, adjacentVertex *vertex.Vertex) {

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
