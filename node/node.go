package node

import (
	"math/big"

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
	// TODO: Add finger table after inefficient solution is
	// complete
}

func New(ipAddr string) *ChordNode {
	return &ChordNode{
		ID:         hash.Hash(ipAddr),
		IpAddr:     ipAddr,
		Successors: nil,
	}
}
