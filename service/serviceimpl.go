package service

import (
	"context"
	"errors"

	"github.com/anthoturc/chord/node"

	pb "github.com/anthoturc/chord/proto"
	"google.golang.org/grpc"
)

type ChordServer struct {
	node   *node.ChordNode
	server *grpc.Server

	// Comment taken from generated gRPC
	// All implementations must embed UnimplementedChordServer
	// for forward compatibility
	pb.UnimplementedChordServer
}

func (s *ChordServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "healthy"}, nil
}

func (s *ChordServer) FindSuccessor(ctx context.Context, req *pb.FindSuccessorRequest) (*pb.FindSuccessorResponse, error) {
	localNode := s.node
	successorAddr, err := localNode.FindSuccessor(req.GetKey())
	return &pb.FindSuccessorResponse{Address: successorAddr}, err
}

func (s *ChordServer) GetPredecessor(ctx context.Context, req *pb.GetPredecessorRequest) (*pb.GetPredecessorResponse, error) {
	localNode := s.node
	if localNode.Predecessor == nil {
		return nil, errors.New("no predecessor for this node")
	}
	return &pb.GetPredecessorResponse{Address: localNode.Predecessor.IpAddr}, nil
}

func (s *ChordServer) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	localNode := s.node
	localNode.Notify(req.GetAddress())

	return &pb.NotifyResponse{}, nil
}
