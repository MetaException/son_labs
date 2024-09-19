package node

func (s Node) Send(dist *Node, count int) *Node {

	if count > len(s.Frames) {
		count = len(s.Frames)
	}

	framesToSend := s.Frames[:count]

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

func (s *Node) DestroyFrames(count int) {
	if count > len(s.Frames) {
		count = len(s.Frames)
	}

	s.Frames = s.Frames[count:]
}

func (node *Node) ClearDeadFramesHistory() {
	for k := range node.FramesIdHistory {
		if node.FramesIdHistory[k] <= 0 {
			delete(node.FramesIdHistory, k)
		}
	}
}
