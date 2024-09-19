package frame

type Frame struct {
	ParentName string
	TTL        int
	ID         string
}

func (f *Frame) String() string {
	return f.ParentName
}
