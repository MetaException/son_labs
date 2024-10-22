package netgraph

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"math"
)

func (g *Graph) ClearMap() {
	for i := range g.VertexMap {
		delete(g.VertexMap, i)
	}
}

func (g *Graph) AddVertex(vertexToAdd vertex.IVertex) {
	if _, ok := g.VertexMap[vertexToAdd]; !ok {
		g.VertexMap[vertexToAdd] = make([]vertex.IVertex, 0)
		g.VertexList = append(g.VertexList, vertexToAdd)

		if n, ok := vertexToAdd.(*node.Node); ok {
			g.Nodes[n.Name] = n
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

	leftNode := ivertexSrc.GetBase()
	rightNode := ivertexToCompare.GetBase()

	return math.Sqrt(math.Pow(rightNode.X-leftNode.X, 2)+math.Pow(rightNode.Y-leftNode.Y, 2)) <= math.Max(rightNode.R, leftNode.R)
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
