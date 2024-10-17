package network

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"os"
	"strconv"
)

type Network struct {
	graph        *Graph
	render       render
	AreaX, AreaY int
}

func NewNetwork(render render, vertexCount int) *Network {
	return &Network{
		render: render,
		graph:  NewGraph(vertexCount),
		AreaX:  100,
		AreaY:  100,
	}
}

func (net *Network) Startup(nodeCount int) {

	fmt.Printf("\nСоздание вершин...\n\n")

	hub := hub.GenerateRandomHub("hub")
	net.graph.AddVertex(hub)

	var lastVertex vertex.Vertex = hub.Vertex
	for i := range nodeCount { // Создаём вершины
		node := node.GenerateRandomNodeByVertex(strconv.Itoa(i+1), lastVertex)
		net.graph.AddVertex(node)

		lastVertex = node.Vertex
	}

	fmt.Printf("\nStarting...\n")

	err := os.RemoveAll("history")
	if err != nil {
		panic(err)
	}
	os.Mkdir("history", 0755)

	net.PerformRounds()
}

func (net *Network) PerformRounds() {
	i := 1
	for !net.graph.CheckAllPoweroff() && !net.graph.CheckFinished() {

		net.graph.ClearMap()
		net.graph.Fill(i)

		if i == 1 {
			if check := net.graph.CheckConnectivity(); !check { // Проверяем граф на связность
				panic("\nСоздан несвязный граф")
			}
		}

		net.PerformRound(i)

		net.render.DrawGraphImage(strconv.Itoa(i), *net.graph)

		net.PerformMoving()

		i++
	}
}

func (net *Network) PerformRound(roundNumber int) {

	for _, sender := range net.graph.Nodes {

		if sender.Power <= 0 {
			continue
		}

		recievers := net.graph.VertexMap[sender]
		for i := range recievers {
			Flooding(sender, recievers[i], sender.FpR)
		}

		sender.DestroyFrames(sender.FpR)
		sender.Power--
		if sender.Power < 0 {
			sender.Power = 0
		}
	}

	net.graph.PrintInfo(roundNumber)
	net.graph.ClearAllDeadFramesHistory()
}

func (net *Network) PerformMoving() {
	for _, node := range net.graph.Nodes {
		node.RandomMove(float64(net.AreaX), float64(net.AreaY))
	}
}
