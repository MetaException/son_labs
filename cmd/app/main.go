package main

import (
	"fmt"
	roundmanager "lab1/internal"
	"lab1/internal/network"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"os"
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
		vertex := lastVertex.GenerateRandomVertexByVertex(strconv.Itoa(i + 1))
		g.AddNode(vertex)

		lastVertex = *vertex
	}

	hub := hub.GenerateRandomHubByBaseNode("hub", lastVertex.Vertex)
	g.AddHub(hub)

	fmt.Printf("\nЗаполнение графа...\n\n")
	g.FillGraph()

	if check := g.CheckConnectivity(); !check { // Проверяем граф на связность
		panic("\nСоздан несвязный граф")
	}

	fmt.Printf("\nStarting...\n")

	err := os.RemoveAll("history")
	if err != nil {
		panic(err)
	}
	os.Mkdir("history", 0755)

	g.PrintInfo(0)
	g.DrawGraph(strconv.Itoa(0))

	r := roundmanager.NewRoundManager(g)

	r.PerformRounds()
}
