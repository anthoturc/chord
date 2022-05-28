package service

//// Lookup API provided by the chord protocol
func (s *ChordServer) Lookup(key string) (string, error) {
	localNode := s.node
	ipAddr, err := localNode.FindSuccessor(key)
	if err != nil {
		return "", err
	}
	return ipAddr, nil
}

//// Access to the underlying ChordNode
func (s *ChordServer) DumpInfo() string {
	return s.node.DumpNodeInfo()
}

func (s *ChordServer) Stop() {
	s.server.GracefulStop()
}
