package vertex

func (v *Vertex) DestroyFrames(count int) {
	if count > len(v.Frames) {
		count = len(v.Frames)
	}

	v.Frames = v.Frames[count:]
}

func (v Vertex) ClearDeadFramesHistory() Vertex {
	for k := range v.FramesIdHistory {
		if v.FramesIdHistory[k] <= 0 {
			delete(v.FramesIdHistory, k)
		}
	}

	return v
}
