package network

import (
	"fmt"
	"lab1/internal/network/netgraph"
	"os"
	"strconv"
)

type Network struct {
	graph  *netgraph.Graph
	render render
}

func NewNetwork(render render, vertexCount int) *Network {
	return &Network{
		render: render,
		graph:  netgraph.NewGraph(vertexCount, 100, 100),
	}
}

func (net *Network) Startup(nodeCount int) {

	fmt.Printf("\nСоздание вершин...\n\n")

	net.graph.GenerateNVertex(nodeCount)

	fmt.Printf("\nStarting...\n")

	err := os.RemoveAll("history")
	if err != nil {
		panic(err)
	}
	os.Mkdir("history", 0755)

	net.graph.Fill(0)

	if check := net.graph.CheckConnectivity(); !check { // Проверяем граф на связность
		panic("\nСоздан несвязный граф")
	}

	net.graph.PrintInfo(0)
	net.render.DrawGraphImage("0", *net.graph)

	net.PerformRounds()
}

func (net *Network) PerformRounds() {
	i := 1
	for !net.graph.CheckAllPoweroff() && !net.graph.CheckFinished() {

		net.graph.PerformRounds(i)

		net.render.DrawGraphImage(strconv.Itoa(i), *net.graph)

		net.graph.PerformMoving()

		i++
	}
}
