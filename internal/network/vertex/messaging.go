package vertex

func (v *Vertex) DestroyFrames(count int) {
	if count > len(v.Frames) {
		count = len(v.Frames)
	}

	v.Frames = v.Frames[count:]
}

func (v Vertex) Send(dist *Vertex, count int) *Vertex {

	if count > len(v.Frames) {
		count = len(v.Frames)
	}

	framesToSend := v.Frames[:count]

	//fmt.Printf("%s %v sends %v to %s %v | ", s.Name, s.Frames, framesToSend, dist.Name, dist.Frames)

	for _, frame := range framesToSend {
		if frame.ParentName != dist.Name {
			if _, ok := dist.FramesIdHistory[frame.ID]; !ok && frame.TTL > 0 {
				frame.TTL--
				dist.FramesIdHistory[frame.ID] = frame.TTL
				dist.Frames = append(dist.Frames, frame)
			}
		}
	}

	//fmt.Printf("result: %v\n", dist.Frames)
	return dist
}

func (v *Vertex) ClearDeadFramesHistory() {
	for k := range v.FramesIdHistory {
		if v.FramesIdHistory[k] <= 0 {
			delete(v.FramesIdHistory, k)
		}
	}
}
