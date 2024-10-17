package network

import (
	"fmt"
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
	var lastVertex node.Node
	for i := range nodeCount { // Создаём вершины
		vertex := lastVertex.GenerateRandomVertexByVertex(strconv.Itoa(i + 1)) // вынести в одну функцию
		net.graph.AddVertex(vertex)

		lastVertex = *vertex
	}

	hub := hub.GenerateRandomHubByBaseNode("hub", lastVertex.Vertex)
	net.graph.AddVertex(hub)

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
		if len(recievers) == 0 {
			continue
		}

		var count int = sender.FpR
		for i := range recievers {
			Flooding(sender, recievers[i], count)
		}

		sender.DestroyFrames(count)
		sender.Power--
	}

	net.graph.PrintInfo(roundNumber)
	net.graph.ClearAllDeadFramesHistory()
}

func (net *Network) PerformMoving() {
	for _, node := range net.graph.Nodes {
		node.RandomMove(float64(net.AreaX), float64(net.AreaY))
	}
}
