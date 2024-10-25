package netgraph

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"strconv"
)

type Graph struct {
	VertexMap  map[vertex.IVertex][]vertex.IVertex // Список смежности
	VertexList []vertex.IVertex                    // Все вершины графа сети

	//RouteMap map[vertex.IVertex]map[vertex.IVertex]*RoutingData
	//RouteList []*RoutingData

	Nodes map[string]*node.Node
	Hubs  map[string]*hub.Hub

	AreaX, AreaY int
}

type RoutingData struct {
	Pintensity float64 // феромоны
	Cost       float64 // hops
}

func NewGraph(N, areaX, areaY int) *Graph {
	return &Graph{
		VertexMap:  make(map[vertex.IVertex][]vertex.IVertex, N),
		Nodes:      make(map[string]*node.Node, 0),
		Hubs:       make(map[string]*hub.Hub),
		VertexList: make([]vertex.IVertex, 0),
		//RouteMap:   make(map[vertex.IVertex]map[vertex.IVertex]*RoutingData),
		//RouteList:  make([]*RoutingData, 0),
		AreaX: areaX,
		AreaY: areaY,
	}
}

func (g *Graph) PrintInfo(roundNumber int) {
	fmt.Println()
	for _, node := range g.VertexList {
		base := node.GetBase()
		fmt.Printf("%v| %s : %v\n", roundNumber, base.Name, base.Frames)
	}
}

func (graph *Graph) GenerateNVertex(nodeCount int) {
	hub := hub.GenerateRandomHub("hub")
	graph.AddVertex(hub)

	var lastVertex vertex.Vertex = hub.Vertex
	for i := range nodeCount { // Создаём вершины
		node := node.GenerateRandomNodeByVertex(strconv.Itoa(i+1), lastVertex)
		graph.AddVertex(node)

		lastVertex = node.Vertex
	}
}
