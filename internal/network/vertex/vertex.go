package vertex

import (
	"lab1/internal/network/frame"
)

type IVertex interface {
	ClearDeadFramesHistory() Vertex
	GetBase() *Vertex
	UpdateBase(base *Vertex) // Метод для обновления базовой структуры
}

type Vertex struct {
	X, Y, R         float64
	Name            string
	Frames          []*frame.Frame
	FramesIdHistory map[string]int
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

func (base Vertex) String() string {
	return base.Name
}

func (base *Vertex) GetBase() *Vertex {
	return base
}
