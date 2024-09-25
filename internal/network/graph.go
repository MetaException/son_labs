package network

import (
	"bytes"
	"fmt"
	"image/png"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/hub"
	"lab1/internal/network/vertex/node"
	"path/filepath"

	"github.com/fogleman/gg"
)

type Graph struct {
	VertexMap    map[*vertex.Vertex][]*vertex.Vertex // Список смежности
	VertexByName map[string]*vertex.Vertex           // Доступ к вершинам по имени
	VertexList   []*vertex.Vertex                    // Все вершины графа сети

	Nodes map[string]*node.Node
	Hubs  map[string]*hub.Hub

	Length       int
	AreaX, AreaY int
}

func NewGraph(N int) *Graph {
	return &Graph{
		VertexMap:    make(map[*vertex.Vertex][]*vertex.Vertex, N),
		Nodes:        make(map[string]*node.Node, 0),
		Length:       N,
		VertexByName: make(map[string]*vertex.Vertex),
		AreaX:        100,
		AreaY:        100,
		Hubs:         make(map[string]*hub.Hub),
		VertexList:   make([]*vertex.Vertex, 0),
	}
}

func (g *Graph) AddHub(hubToAdd *hub.Hub) {
	if _, ok := g.VertexMap[&hubToAdd.Vertex]; !ok {
		g.VertexMap[&hubToAdd.Vertex] = make([]*vertex.Vertex, 0)
		g.Hubs[hubToAdd.Vertex.Name] = hubToAdd
		g.VertexList = append(g.VertexList, &hubToAdd.Vertex)
	}
}

func (g *Graph) AddNode(nodeToAdd *node.Node) {

	if _, ok := g.VertexMap[&nodeToAdd.Vertex]; !ok {
		g.VertexMap[&nodeToAdd.Vertex] = make([]*vertex.Vertex, 0)
		g.Nodes[nodeToAdd.Name] = nodeToAdd
		g.VertexByName[nodeToAdd.Name] = &nodeToAdd.Vertex
		g.VertexList = append(g.VertexList, &nodeToAdd.Vertex)
	}
}

func (g *Graph) AddEdge(vertex *vertex.Vertex, adjacentVertex *vertex.Vertex) {

	g.VertexMap[vertex] = append(g.VertexMap[vertex], adjacentVertex)
	g.VertexMap[adjacentVertex] = append(g.VertexMap[adjacentVertex], vertex)
}

func (g *Graph) FillGraph() {
	for i := 0; i < len(g.VertexList); i++ {
		for j := i + 1; j < len(g.VertexList); j++ {

			vertex := g.VertexList[i]
			vertexToCompare := g.VertexList[j]

			if vertex.IsAdjacent(vertexToCompare) {
				g.AddEdge(vertex, vertexToCompare)
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
	for startNode, neighbors := range g.VertexMap {
		for _, endNode := range neighbors {
			dc.DrawLine(startNode.X*scale+padding, startNode.Y*scale+padding, endNode.X*scale+padding, endNode.Y*scale+padding)
			dc.Stroke()
		}
	}

	// Рисуем узлы (в виде окружностей) и их имена
	for _, node := range g.Nodes {
		drawNode(dc, node, scale, padding)
	}

	for _, hub := range g.Hubs {
		drawHub(dc, hub, scale, padding)
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

func drawHub(dc *gg.Context, node *hub.Hub, scale, padding float64) {
	dc.SetRGBA(0, 0, 0, 0.5) // Цвет радиуса (черный)
	dc.SetLineWidth(1)
	dc.DrawCircle(node.X*scale+padding, node.Y*scale+padding, node.R*scale)
	dc.Stroke()

	dc.DrawCircle(node.X*scale+padding, node.Y*scale+padding, 5)
	dc.SetRGB(1, 0, 0)

	dc.Fill()

	// Добавляем название узлов
	dc.SetRGB(0, 0, 0) // Цвет текста (черный)

	dc.DrawStringAnchored(node.Name, node.X*scale+padding, node.Y*scale+padding-64/2-32, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("(%.0f;%.0f)", node.X, node.Y), node.X*scale+padding, node.Y*scale+padding-64/2-16, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("R: %v", int(node.R)), node.X*scale+padding, node.Y*scale+padding-64/2, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("Frames: %v", len(node.Frames)), node.X*scale+padding, node.Y*scale+padding-64/2+16, 0.5, 0.5)
	dc.Stroke()
}

func drawNode(dc *gg.Context, node *node.Node, scale, padding float64) {
	dc.SetRGBA(0, 0, 0, 0.5) // Цвет радиуса (черный)
	dc.SetLineWidth(1)
	dc.DrawCircle(node.X*scale+padding, node.Y*scale+padding, node.R*scale)
	dc.Stroke()

	dc.DrawCircle(node.X*scale+padding, node.Y*scale+padding, 5)
	dc.SetRGB(0, 0.5, 0.8) // Цвет узлов (голубой)
	dc.Fill()

	// Добавляем название узлов
	dc.SetRGB(0, 0, 0) // Цвет текста (черный)

	dc.DrawStringAnchored(node.Name, node.X*scale+padding, node.Y*scale+padding-64/2-32, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("(%.0f;%.0f)", node.X, node.Y), node.X*scale+padding, node.Y*scale+padding-64/2-16, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("R: %v", int(node.R)), node.X*scale+padding, node.Y*scale+padding-64/2, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("FpR: %v; Power: %.2f; Frames: %v", node.FpR, node.Power, len(node.Frames)), node.X*scale+padding, node.Y*scale+padding-64/2+16, 0.5, 0.5)
	dc.Stroke()
}
