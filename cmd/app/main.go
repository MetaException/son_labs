package main

import (
	"fmt"
	"lab1/internal/network"
	"lab1/internal/node"
	"lab1/internal/roundmanager"
	"strconv"
)

func main() {

	var nodeCount int
	fmt.Print("Введите количество вершин: ")
	fmt.Scanln(&nodeCount)

	g := network.NewGraph(nodeCount)

	fmt.Printf("\nСоздание вершин...\n\n")
	var lastVertex node.Node
	for i := range g.Length { // Создаём вершины
		vertex := lastVertex.GenerateRandomVertexByVertex(strconv.Itoa(i+1), false)
		g.AddNode(vertex)

		lastVertex = *vertex
	}

	hub := lastVertex.GenerateRandomVertexByVertex("hub", true)
	g.AddNode(hub)

	fmt.Printf("\nЗаполнение графа...\n\n")
	g.FillGraph()

	if check := g.CheckConnectivity(); !check { // Проверяем граф на связность
		panic("\nСоздан несвязный граф")
	}

	fmt.Printf("\nStarting...\n")

	g.PrintInfo(0)

	r := roundmanager.NewRoundManager(g)

	r.PerformRounds()
}
