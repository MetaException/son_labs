package network

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
)

type Graph struct {
	VertexMap  map[vertex.IVertex][]vertex.IVertex // Список смежности
	VertexList []vertex.IVertex                    // Все вершины графа сети

	VertexByCluster    map[int][]*node.Node
	ClusterHeadHistory map[*node.Node]struct{}
	CurrentHeadList    map[*node.Node]struct{}

	Nodes map[string]*node.Node
	Hubs  map[string]*hub.Hub
}

func NewGraph(N int) *Graph {
	return &Graph{
		VertexMap:          make(map[vertex.IVertex][]vertex.IVertex, N),
		Nodes:              make(map[string]*node.Node, 0),
		Hubs:               make(map[string]*hub.Hub),
		VertexList:         make([]vertex.IVertex, 0),
		VertexByCluster:    make(map[int][]*node.Node),
		ClusterHeadHistory: make(map[*node.Node]struct{}),
		CurrentHeadList:    make(map[*node.Node]struct{}),
	}
}

func (g *Graph) ClearMap() {
	for i := range g.VertexMap {
		delete(g.VertexMap, i)
	}
}

func (g *Graph) ClearHeadHistory() {

	if len(g.ClusterHeadHistory) == len(g.Nodes) {
		for n := range g.ClusterHeadHistory {
			delete(g.ClusterHeadHistory, n)
		}
	}

	for n := range g.CurrentHeadList {
		delete(g.CurrentHeadList, n)
	}
}

func (g *Graph) AddVertex(vertexToAdd vertex.IVertex) {
	if _, ok := g.VertexMap[vertexToAdd]; !ok {
		g.VertexMap[vertexToAdd] = make([]vertex.IVertex, 0)
		g.VertexList = append(g.VertexList, vertexToAdd)

		if n, ok := vertexToAdd.(*node.Node); ok {
			g.Nodes[n.Name] = n

			g.VertexByCluster[n.Cluster] = append(g.VertexByCluster[n.Cluster], n)

		} else if hub, ok := vertexToAdd.(*hub.Hub); ok {
			g.Hubs[hub.Name] = hub
		}
	}
}

func (g *Graph) AddEdge(vertex vertex.IVertex, adjacentVertex vertex.IVertex) {

	g.VertexMap[vertex] = append(g.VertexMap[vertex], adjacentVertex)
	g.VertexMap[adjacentVertex] = append(g.VertexMap[adjacentVertex], vertex)
}

func (g *Graph) Fill(roundNumber int) {
	for i := 0; i < len(g.VertexList); i++ {
		for j := i + 1; j < len(g.VertexList); j++ {

			v := g.VertexList[i]
			vertexToCompare := g.VertexList[j]

			if g.IsAdjacent(v, vertexToCompare) {
				g.AddEdge(v, vertexToCompare)
			}
		}
	}

	fmt.Println(g.VertexMap)
}

func (g *Graph) IsAdjacent(ivertexSrc vertex.IVertex, ivertexToCompare vertex.IVertex) bool {

	leftNode, okleft := ivertexSrc.(*node.Node)
	rightNode, okright := ivertexToCompare.(*node.Node)

	_, isRightNodeHead := g.CurrentHeadList[rightNode]
	_, isLeftNodeHead := g.CurrentHeadList[leftNode]

	if !okleft {
		return isRightNodeHead
	}
	if !okright {
		return isLeftNodeHead
	}

	if isLeftNodeHead && isRightNodeHead {
		return false
	}

	return leftNode.Cluster == rightNode.Cluster && (isLeftNodeHead || isRightNodeHead)
}

func (g *Graph) PrintInfo(roundNumber int) {
	fmt.Println()
	for _, node := range g.VertexList {
		base := node.GetBase()
		fmt.Printf("%v| %s : %v\n", roundNumber, base.Name, base.Frames)
	}
}

func (g *Graph) CheckConnectivity() bool {

	if len(g.VertexMap) == 0 {
		return true
	}

	isFirst := false
	var first vertex.IVertex

	hashset := make(map[vertex.IVertex]struct{})

	for _, node := range g.VertexList {
		if !isFirst {
			first = node
			isFirst = true
		}
		hashset[node] = struct{}{}
	}

	queue := []vertex.IVertex{first}
	delete(hashset, first)

	for len(queue) > 0 {

		qNode := queue[0]
		queue = queue[1:]
		for _, node := range g.VertexMap[qNode] {
			if _, ok := hashset[node]; ok {
				queue = append(queue, node)
				delete(hashset, node)
			}
		}

		if len(hashset) == 0 {
			return true
		}
	}
	fmt.Println(hashset)
	return false
}
