package network

import (
	"fmt"
	"lab1/internal/network/netgraph"
	"lab1/internal/network/vertex"
	"lab1/internal/network/vertex/node"
	"path/filepath"

	"github.com/fogleman/gg"
)

const width = 1000
const height = 1000
const fieldSize = 100.0
const padding = 100.0
const scale = (width - 2*padding) / fieldSize

type render struct {
	gg *gg.Context
}

func NewRender() *render {
	return &render{
		gg: gg.NewContext(width, height),
	}
}

func (r render) DrawGraphImage(name string, graph netgraph.Graph) {

	// Заливаем фон белым цветом
	r.gg.SetRGB(1, 1, 1)
	r.gg.Clear()

	r.drawVertexes(graph.VertexList)
	r.drawEdges(graph.VertexMap)
	r.drawRoutes(graph.RouteMap)

	err := r.gg.SavePNG(filepath.Join("history", name+".png"))
	if err != nil {
		panic(err)
	}
}

func (r render) drawRoutes(graph map[vertex.IVertex]map[vertex.IVertex]*netgraph.RoutingData) {
	for vertex, edges := range graph {
		r.drawRoute(vertex, edges)
	}
}

func (r render) drawRoute(vertex vertex.IVertex, edges map[vertex.IVertex]*netgraph.RoutingData) {
	baseVertex := vertex.GetBase()
	r.gg.SetLineWidth(3.0)
	fmt.Println()
	fmt.Print(vertex, " ")

	maxVal := 0.0
	for _, edge := range edges {
		if edge.Pintensity > float64(maxVal) {
			maxVal = edge.Pintensity
		}
	}

	for v, edge := range edges {
		color := edge.Pintensity
		if edge.Pintensity > 1 {
			color = edge.Pintensity / maxVal
		}
		r.gg.SetRGBA(1, 0, 0, color) // TODO нормализовать значения
		baseEdge := v.GetBase()
		r.gg.DrawLine(baseVertex.X*scale+padding, baseVertex.Y*scale+padding, baseEdge.X*scale+padding, baseEdge.Y*scale+padding)
		r.gg.Stroke()
	}
}

func (r render) drawVertexes(list []vertex.IVertex) {
	for _, vertex := range list {
		r.drawVertex(vertex)
	}
}

func (r render) drawEdges(graph map[vertex.IVertex][]vertex.IVertex) {
	r.gg.SetRGBA(0, 0, 0, 0.5) // Цвет текста (черный)
	for vertex, edges := range graph {
		r.drawEdge(vertex, edges)
	}
}

func (r render) drawEdge(vertex vertex.IVertex, edges []vertex.IVertex) {
	baseVertex := vertex.GetBase()
	for _, edge := range edges {
		baseEdge := edge.GetBase()
		r.gg.DrawLine(baseVertex.X*scale+padding, baseVertex.Y*scale+padding, baseEdge.X*scale+padding, baseEdge.Y*scale+padding)
		r.gg.Stroke()
	}
}

func (r render) drawVertex(vertex vertex.IVertex) {

	vbase := vertex.GetBase()

	r.gg.DrawCircle(vbase.X*scale+padding, vbase.Y*scale+padding, 5)

	if _, ok := vertex.(*node.Node); ok {
		r.gg.SetRGB(0, 0.5, 0.8) // Цвет узлов (голубой)
	} else {
		r.gg.SetRGB(1, 0, 0)
	}

	r.gg.Fill()
	r.gg.SetRGB(0, 0, 0) // Цвет текста (черный)

	r.gg.DrawStringAnchored(vbase.Name, vbase.X*scale+padding, vbase.Y*scale+padding-64/2-16, 0.5, 0.5)

	if node, ok := vertex.(*node.Node); ok {
		r.gg.DrawStringAnchored(fmt.Sprintf("P: %.2f; F: %v", node.Power, len(node.Frames)), vbase.X*scale+padding, vbase.Y*scale+padding-64/2, 0.5, 0.5)
	} else {
		r.gg.DrawStringAnchored(fmt.Sprintf("F: %v", len(vbase.Frames)), vbase.X*scale+padding, vbase.Y*scale+padding-64/2, 0.5, 0.5)
	}
	r.gg.Stroke()

	r.gg.SetRGBA(0, 0, 0, 0.5) // Цвет радиуса (черный)
	r.gg.SetLineWidth(1)
	r.gg.DrawCircle(vbase.X*scale+padding, vbase.Y*scale+padding, vbase.R*scale)
	r.gg.Stroke()
}
