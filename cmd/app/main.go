package main

import (
	"fmt"
	"lab1/internal/network"
)

func main() {

	var nodeCount int
	fmt.Print("Введите количество вершин: ")
	fmt.Scanln(&nodeCount)

	render := network.NewRender()
	net := network.NewNetwork(*render, nodeCount)

	net.Startup(nodeCount)
}
