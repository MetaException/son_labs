package network

import (
	"fmt"
	"lab1/internal/network/vertex"
)

func (g *Graph) PrintInfo(roundNumber int) {
	fmt.Println()
	for _, node := range g.VertexList {
		fmt.Printf("%v| %s : %v\n", roundNumber, node.Name, node.Frames)
	}
}

func (g *Graph) CheckConnectivity() bool {

	if len(g.VertexMap) == 0 {
		return true
	}

	isFirst := false
	var first *vertex.Vertex

	hashset := make(map[*vertex.Vertex]struct{})

	vertexList := make([]*vertex.Vertex, 0)

	for _, node := range g.Nodes {
		vertexList = append(vertexList, &node.Vertex)
	}

	for _, hub := range g.Hubs {
		vertexList = append(vertexList, &hub.Vertex)
	}

	for _, node := range vertexList {
		if !isFirst {
			first = node
			isFirst = true
		}
		hashset[node] = struct{}{}
	}

	queue := []*vertex.Vertex{first}
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
