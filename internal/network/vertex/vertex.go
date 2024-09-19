package vertex

import "lab1/internal/network/frame"

type Vertex struct {
	X, Y, R         float64
	Name            string
	Frames          []*frame.Frame // Сделать лимит??
	FramesIdHistory map[string]int
}

func (base Vertex) String() string {
	return base.Name
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
