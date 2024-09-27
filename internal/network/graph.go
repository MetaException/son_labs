package network

import (
	"bytes"
	"fmt"
	"image/png"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"math"
	"math/rand"
	"path/filepath"

	"github.com/fogleman/gg"
)

type Graph struct {
	VertexMap       map[vertex.IVertex][]vertex.IVertex // Список смежности
	VertexList      []vertex.IVertex                    // Все вершины графа сети
	G               map[*node.Node]struct{}
	VertexByCluster map[int][]*node.Node

	Nodes map[string]*node.Node
	Hubs  map[string]*hub.Hub

	Length       int
	AreaX, AreaY int
}

func NewGraph(N int) *Graph {
	return &Graph{
		VertexMap:       make(map[vertex.IVertex][]vertex.IVertex, N),
		Nodes:           make(map[string]*node.Node, 0),
		Length:          N,
		AreaX:           100,
		AreaY:           100,
		Hubs:            make(map[string]*hub.Hub),
		VertexList:      make([]vertex.IVertex, 0),
		G:               make(map[*node.Node]struct{}),
		VertexByCluster: make(map[int][]*node.Node),
	}
}

func (g *Graph) AddVertex(vertexToAdd vertex.IVertex) {
	if _, ok := g.VertexMap[vertexToAdd]; !ok {
		g.VertexMap[vertexToAdd] = make([]vertex.IVertex, 0)
		g.VertexList = append(g.VertexList, vertexToAdd)

		if n, ok := vertexToAdd.(*node.Node); ok {
			g.Nodes[n.Name] = n

			g.VertexByCluster[n.Cluster] = append(g.VertexByCluster[n.Cluster], n)

		} else if hub, ok := vertexToAdd.(*hub.Hub); ok {
			g.Hubs[hub.Name] = hub
		}
	}
}

func (g *Graph) AddEdge(vertex vertex.IVertex, adjacentVertex vertex.IVertex) {

	g.VertexMap[vertex] = append(g.VertexMap[vertex], adjacentVertex)
	g.VertexMap[adjacentVertex] = append(g.VertexMap[adjacentVertex], vertex)
}

func (g *Graph) CalculateTn(r int) {

	CHCandidates := make(map[int][]*node.Node)

	for _, node := range g.Nodes {

		P := 0.04 //1.0 / float64(len(g.VertexByCluster[node.Cluster]))

		var tn float64
		if _, ok := g.G[node]; !ok {
			tn = P / (1 - P*math.Mod(float64(r), 1/P))
		} else {
			tn = 0
		}

		if rand.Float64() < tn {
			CHCandidates[node.Cluster] = append(CHCandidates[node.Cluster], node)
		}
	}

	for k, v := range CHCandidates {
		fmt.Printf("for %v: %v\n", k, v)
		if len(v) != 1 {
			g.Nodes[CHCandidates[k][0].Name].IsClusterHead = true // Для примера берём первый можно вынести в отденый список
			g.G[CHCandidates[k][0]] = struct{}{}
		} else { // Выбираем по энергии и по расстоянию

		}
	}
}

func (g *Graph) FillGraph(roundNumber int) {

	g.CalculateTn(roundNumber)
	for i := 0; i < len(g.VertexList); i++ {
		for j := i + 1; j < len(g.VertexList); j++ {

			v := g.VertexList[i]
			vertexToCompare := g.VertexList[j]

			if IsAdjacent(v, vertexToCompare) {
				g.AddEdge(v, vertexToCompare)
			}
		}
	}

	fmt.Println(g.VertexMap)
}

func (g *Graph) DrawGraph(name string) {
	const width = 1000
	const height = 1000
	const fieldSize = 100.0
	const padding = 100.0
	scale := (width - 2*padding) / fieldSize

	// Создаем новое изображение
	dc := gg.NewContext(width, height)

	// Заливаем фон белым цветом
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Рисуем рёбра (линии между смежными узлами)
	dc.SetRGB(0, 0, 0) // Цвет линии (черный)
	for istartNode, neighbors := range g.VertexMap {
		startNode := istartNode.GetBase()
		for _, iendNode := range neighbors {
			endNode := iendNode.GetBase()
			dc.DrawLine(startNode.X*scale+padding, startNode.Y*scale+padding, endNode.X*scale+padding, endNode.Y*scale+padding)
			dc.Stroke()
		}
	}

	for _, node := range g.VertexList {
		drawVertex(dc, node, scale, padding)
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, dc.Image()); err != nil {
		panic(err)
	}

	err := dc.SavePNG(filepath.Join("history", name+".png"))
	if err != nil {
		panic(err)
	}
}

func IsAdjacent(ivertexSrc vertex.IVertex, ivertexToCompare vertex.IVertex) bool {

	leftNode, okleft := ivertexSrc.(*node.Node)
	rightNode, okright := ivertexToCompare.(*node.Node)

	if !okleft {
		return rightNode.IsClusterHead
	}
	if !okright {
		return leftNode.IsClusterHead
	}

	if leftNode.IsClusterHead && rightNode.IsClusterHead {
		return false
	}

	return leftNode.Cluster == rightNode.Cluster && (leftNode.IsClusterHead || rightNode.IsClusterHead)
}

func drawVertex(dc *gg.Context, vertex vertex.IVertex, scale, padding float64) {

	vbase := vertex.GetBase()

	//dc.SetRGBA(0, 0, 0, 0.5) // Цвет радиуса (черный)
	//dc.SetLineWidth(1)
	//dc.DrawCircle(vbase.X*scale+padding, vbase.Y*scale+padding, vbase.R*scale)
	//dc.Stroke()

	dc.DrawCircle(vbase.X*scale+padding, vbase.Y*scale+padding, 5)

	if _, ok := vertex.(*node.Node); ok {
		dc.SetRGB(0, 0.5, 0.8) // Цвет узлов (голубой)
	} else {
		dc.SetRGB(1, 0, 0)
	}

	dc.Fill()

	// Добавляем название узлов
	dc.SetRGB(0, 0, 0) // Цвет текста (черный)

	dc.DrawStringAnchored(vbase.Name, vbase.X*scale+padding, vbase.Y*scale+padding-64/2-32, 0.5, 0.5)
	//dc.DrawStringAnchored(fmt.Sprintf("(%.0f;%.0f)", vbase.X, vbase.Y), vbase.X*scale+padding, vbase.Y*scale+padding-64/2-16, 0.5, 0.5)
	//dc.DrawStringAnchored(fmt.Sprintf("R: %.0f", vbase.R), vbase.X*scale+padding, vbase.Y*scale+padding-64/2, 0.5, 0.5)

	if node, ok := vertex.(*node.Node); ok {
		dc.DrawStringAnchored(fmt.Sprintf("Cluster: %v", node.Cluster), vbase.X*scale+padding, vbase.Y*scale+padding-64/2+16, 0.5, 0.5)
		//dc.DrawStringAnchored(fmt.Sprintf("FpR: %v; Power: %.2f; Frames: %v; Cluster: %v", node.FpR, node.Power, len(node.Frames), node.Cluster), vbase.X*scale+padding, vbase.Y*scale+padding-64/2+16, 0.5, 0.5)
	} else {
		dc.DrawStringAnchored(fmt.Sprintf("Frames: %v", len(vbase.Frames)), vbase.X*scale+padding, vbase.Y*scale+padding-64/2+16, 0.5, 0.5)
	}
	dc.Stroke()
}
