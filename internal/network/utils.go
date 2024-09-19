package network

import (
	"fmt"
	"lab1/internal/node"
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
	var first *node.Node

	hashset := make(map[*node.Node]struct{})
	for _, node := range g.VertexList {
		if !isFirst {
			first = node
			isFirst = true
		}
		hashset[node] = struct{}{}
	}

	queue := []*node.Node{first}
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
