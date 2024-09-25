package network

import (
	"fmt"
	"lab1/internal/network/vertex"
)

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
