package node

import (
	"math/big"

	"github.com/anthoturc/chord/client"
	"github.com/anthoturc/chord/hash"
)

// Represents the chord node as described
// in the Chord paper.
type ChordNode struct {
	// Store the hash of the IpAddr to avoid recomputing the
	// ID. Big int is being used to allow for operations on the 160-bit IDs
	// that are referenced in the paper
	ID *big.Int
	// Each node must be aware of its own address + port
	// The expected format of this field is IP:PORT
	IpAddr string

	Predecessor *ChordNode
	Successors  []*ChordNode
}

func New(ipAddr string) *ChordNode {
	return &ChordNode{
		ID:     hash.Hash(ipAddr),
		IpAddr: ipAddr,
		// Using 1-based indexing and temporarily only using 1 successor
		Successors:  make([]*ChordNode, 2),
		Predecessor: nil,
	}
}

// Methods from Figure 6 of Chord paper
func (n *ChordNode) Join(remoteAddr string) error {
	successorAddr, err := client.CallFindSuccessor(remoteAddr, n.IpAddr)
	if err != nil {
		return err
	}

	n.Successors[1] = New(successorAddr)

	client.CallNotify(successorAddr, n.IpAddr)

	return nil
}

func (n *ChordNode) FindSuccessor(key string) (string, error) {
	localSuccessor := n.Successors[1] // Potentially invalid information since this successor may not exist later
	id := hash.Hash(key)
	if hash.IsBetween(n.ID, id, localSuccessor.ID, true) {
		return localSuccessor.IpAddr, nil
	}

	// TODO: Make use of closest_preceding_node (and finger tables) as defined in Chord paper
	remoteSuccessorIpAddr, err := client.CallFindSuccessor(localSuccessor.IpAddr, key)
	if err != nil {
		return "", err
	}

	return remoteSuccessorIpAddr, nil
}

// The predecessor think's it could be n's predecessor
func (n *ChordNode) Notify(predecessorAddr string) {
	maybePredecessorNode := New(predecessorAddr)
	if n.Predecessor == nil || hash.IsBetween(n.Predecessor.ID, maybePredecessorNode.ID, n.ID, false) {
		n.Predecessor = maybePredecessorNode
	}
}

// Run periodically to refresh the finger table entries
func (n *ChordNode) FixFingers() {
	successor, err := n.FindSuccessor(n.IpAddr)
	if err != nil {
		n.Successors[1] = n
		return
	}

	n.Successors[1] = New(successor)
}

func (n *ChordNode) CheckPredecessor() {
	if n.Predecessor == nil {
		return
	}

	msg, err := client.CallPing(n.Predecessor.IpAddr)
	if err != nil || msg != "healthy" {
		n.Predecessor = nil
	}
}

// Run periodically to verify this node's immediate successor and
// tells the successor about n
func (n *ChordNode) Stabilize() {
	localSuccessor := n.Successors[1]

	remotePredecessorIp, err := client.CallGetPredecessor(localSuccessor.IpAddr)
	if err != nil {
		return
	}

	x := New(remotePredecessorIp)

	if hash.IsBetween(n.ID, x.ID, localSuccessor.ID, false) {
		n.Successors[1] = x
	}

	client.CallNotify(n.Successors[1].IpAddr, n.IpAddr)
}

func (n *ChordNode) DumpNodeInfo() string {

	successorIpAddr := "none"
	successorId := "none"
	localSuccessor := n.Successors[1]
	if localSuccessor != nil {
		successorId = localSuccessor.ID.String()
		successorIpAddr = localSuccessor.IpAddr
	}

	predecessorIpAddr := "none"
	predecessorId := "none"
	localPredecessor := n.Predecessor
	if localPredecessor != nil {
		predecessorId = localPredecessor.ID.String()
		predecessorIpAddr = localPredecessor.IpAddr
	}

	return "{\n" +
		"  \"Id\": " + n.ID.String() + ",\n" +
		"  \"IpAddress:\" " + n.IpAddr + ",\n" +
		"  \"SuccessorId\": " + successorId + ",\n" +
		"  \"SuccessorIpAddr\": " + successorIpAddr + ",\n" +
		"  \"PredecessorId\": " + predecessorId + ",\n" +
		"  \"PredecessorIpAddr\": " + predecessorIpAddr + "\n" +
		"}\n"

}
