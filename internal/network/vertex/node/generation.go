package node

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/pkg/utils"
)

func GenerateRandomNode(name string) *Node {
	base := vertex.GenerateRandomBase(name)
	nodeFrameCount := utils.GenerateRandomInt(5, 10)
	fpr := utils.GenerateRandomInt(1, 5)

	fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, base.X, base.Y, base.R, nodeFrameCount)

	return NewNode(base, fpr, nodeFrameCount)
}

func GenerateRandomNodeByVertex(name string, source vertex.Vertex) *Node {
	base := vertex.GenerateRandomBaseByVertex(name, source)
	nodeFrameCount := utils.GenerateRandomInt(2, 7)
	fpr := utils.GenerateRandomInt(1, 5)

	fmt.Printf("New vertex [%s] : X: %v, Y: %v, R: %v, FC: %v\n", name, base.X, base.Y, base.R, nodeFrameCount)

	return NewNode(base, fpr, nodeFrameCount)
}
