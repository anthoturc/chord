package chord

import (
	"context"

	pb "github.com/anthoturc/chord/proto"
)

type ChordServer struct {
	// TODO: Add fields that the server will need here
	// e.g this could be data base connections

	// Comment taken from generated gRPC
	// All implementations must embed UnimplementedChordServer
	// for forward compatibility
	pb.UnimplementedChordServer
}

func (s *ChordServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "healthy"}, nil
}

func (s *ChordServer) FindSuccessor(ctx context.Context, req *pb.FindSuccessorRequest) (*pb.FindSuccessorResponse, error) {
	return &pb.FindSuccessorResponse{Address: "0.0.0.0:1234"}, nil
}

func (s *ChordServer) GetPredecessor(ctx context.Context, req *pb.GetPredecessorRequest) (*pb.GetPredecessorResponse, error) {
	return &pb.GetPredecessorResponse{Address: "0.0.0.0:1234"}, nil
}
