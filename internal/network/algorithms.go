package network

import (
	"fmt"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/node"
	"math"
	"math/rand"
)

func Flooding(src *node.Node, dist vertex.IVertex, count int) { // TODO: убрать в другое место??

	if count > len(src.Vertex.Frames) {
		count = len(src.Vertex.Frames)
	}

	framesToSend := src.Vertex.Frames[:count]

	sentCount := 0.0
	reciever := dist.GetBase()

	for _, frame := range framesToSend {
		if frame.ParentName != reciever.Name {
			if _, ok := reciever.FramesIdHistory[frame.ID]; !ok && frame.TTL > 0 {
				reciever.FramesIdHistory[frame.ID] = frame.TTL
				reciever.Frames = append(reciever.Frames, frame)
				frame.TTL--
				sentCount++
			}
		}
	}

	//	fmt.Printf("\n%s %v sends %v to %s %v", src.Name, src.Frames, framesToSend, reciever.Name, reciever.Frames)

	if node, ok := dist.(*node.Node); ok {
		node.Power -= 0.2 * float64(sentCount)
	}
}

func (graph *Graph) CalculateTn(r int) {

	CHCandidates := make(map[int][]*node.Node)

	for _, node := range graph.Nodes {

		P := 0.04 //1.0 / float64(len(g.VertexByCluster[node.Cluster]))

		var tn float64
		if _, ok := graph.ClusterHeadHistory[node]; !ok {
			tn = P / (1 - P*math.Mod(float64(r), 1/P))
		} else {
			tn = 0
		}

		if rand.Float64() < tn {
			CHCandidates[node.Cluster] = append(CHCandidates[node.Cluster], node)
		}
	}

	for k := 1; k <= 4; k++ {

		v, ok := CHCandidates[k]

		fmt.Printf("for %v: %v\n", k, v)
		if ok {
			graph.ClusterHeadHistory[v[0]] = struct{}{}
			graph.CurrentHeadList[v[0]] = struct{}{}
		} else {
			// Если ни один не попал, то выбираем первый попавшийся для конкретного кластера
			//fmt.Println("hello")
			//fmt.Println(graph.VertexByCluster)

			isFound := graph.pickupRandomNode(k)

			// Если уже прям все узлы были хотябы раз, то берём первый в списке
			if !isFound {
				graph.ClusterHeadHistory[graph.VertexByCluster[k][0]] = struct{}{}
				graph.CurrentHeadList[graph.VertexByCluster[k][0]] = struct{}{}
			}
		}
	}
}

func (graph *Graph) pickupRandomNode(k int) bool {
	for _, v := range graph.VertexByCluster[k] {
		if _, ok := graph.ClusterHeadHistory[v]; !ok {
			graph.ClusterHeadHistory[v] = struct{}{}
			graph.CurrentHeadList[v] = struct{}{}
			return true
		}
	}
	return false
}

func (graph *Graph) ClearAllDeadFramesHistory() {
	for _, vertex := range graph.VertexList {
		vertex.ClearDeadFramesHistory()
	}
}

func (graph *Graph) CheckFinished() bool {

	for _, node := range graph.Nodes {
		if len(node.Frames) != 0 {
			return false
		}
	}
	return true
}

func (graph *Graph) CheckAllPoweroff() bool {

	for _, node := range graph.Nodes {
		if node.Power > 0 {
			return false
		}
	}
	return true
}
