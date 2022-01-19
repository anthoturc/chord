package service

import (
	"context"
	"testing"

	pb "github.com/anthoturc/chord/proto"
)

// TODO: For the APIs, try to figure out how to use a testing table
// since (for now) the "assertions" are pretty similar
func TestFindSuccessor(t *testing.T) {
	s := &ChordServer{}

	resp, err := s.FindSuccessor(context.Background(), &pb.FindSuccessorRequest{Key: "someid"})
	if err != nil {
		t.Errorf("FindSuccessor returned unexpected err: %v\n", err)
	}

	// TODO: WHen this API is implemented we will want to verify
	// that the address is not nil, that is matches an IP:PORT format
	if resp.GetAddress() != "0.0.0.0:1234" {
		t.Errorf("Wanted \"0.0.0.0:1234\" but got %s", resp.GetAddress())
	}
}

func TestGetPredecessor(t *testing.T) {
	s := &ChordServer{}

	resp, err := s.GetPredecessor(context.Background(), &pb.GetPredecessorRequest{})
	if err != nil {
		t.Errorf("FindSuccessor returned unexpected err: %v\n", err)
	}

	if resp.GetAddress() != "0.0.0.0:1234" {
		t.Errorf("Wanted \"0.0.0.0:1234\" but got %s", resp.GetAddress())
	}
}

func TestPingServer(t *testing.T) {
	s := &ChordServer{}

	resp, err := s.Ping(context.Background(), &pb.PingRequest{})
	if err != nil {
		t.Errorf("Ping returned unexpected err: %v\n", err)
	}

	if resp.GetMessage() != "healthy" {
		t.Errorf("Wanted \"healthy\" but got %s", resp.GetMessage())
	}
}
