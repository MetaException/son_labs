package vertex

import (
	"lab1/internal/network/frame"
	"lab1/pkg/utils"
	"math"
)

type Vertex struct {
	X, Y, R         float64
	Name            string
	Frames          []*frame.Frame // Сделать лимит??
	FramesIdHistory map[string]int
}

type IVertex interface {
	ClearDeadFramesHistory() Vertex
	GetBase() *Vertex
	UpdateBase(base *Vertex) // Метод для обновления базовой структуры
}

func (base Vertex) String() string {
	return base.Name
}

func (base *Vertex) GetBase() *Vertex {
	return base
}

func NewBaseNode(X, Y, R float64, Name string) *Vertex {
	return &Vertex{
		X:               X,
		Y:               Y,
		R:               R,
		Name:            Name,
		Frames:          make([]*frame.Frame, 0),
		FramesIdHistory: make(map[string]int),
	}
}

func GenerateRandomBase(name string) *Vertex {
	x := float64(utils.GenerateRandomInt(0, 100))
	y := float64(utils.GenerateRandomInt(0, 100))
	r := float64(utils.GenerateRandomInt(5, 50))

	return NewBaseNode(float64(int(x)), float64(int(y)), r, name)
}

func GenerateRandomBaseByVertex(name string, source Vertex) *Vertex {

	base := GenerateRandomBase(name)

	cx := source.X
	cy := source.Y

	dx := base.X - cx
	dy := base.Y - cy
	distance := math.Sqrt(dx*dx + dy*dy)

	cr := base.R

	if distance > cr {
		ratio := cr / distance
		base.X = cx + dx*ratio
		base.Y = cy + dy*ratio
	}

	return base
}
