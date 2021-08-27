package chord

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/anthoturc/chord/node"
	pb "github.com/anthoturc/chord/proto"
	"google.golang.org/grpc"
)

type ChordServer struct {
	// TODO: Add fields that the server will need here
	// e.g this could be data base connections
	Node *pb.ChordServer

	// Comment taken from generated gRPC
	// All implementations must embed UnimplementedChordServer
	// for forward compatibility
	pb.UnimplementedChordServer
}

//// Chord gRPC API definitions
func (s *ChordServer) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "healthy"}, nil
}

func (s *ChordServer) FindSuccessor(ctx context.Context, req *pb.FindSuccessorRequest) (*pb.FindSuccessorResponse, error) {
	return &pb.FindSuccessorResponse{Address: "0.0.0.0:1234"}, nil
}

func (s *ChordServer) GetPredecessor(ctx context.Context, req *pb.GetPredecessorRequest) (*pb.GetPredecessorResponse, error) {
	return &pb.GetPredecessorResponse{Address: "0.0.0.0:1234"}, nil
}

////

//// Chord service initialization
func Init() (*node.ChordNode, *grpc.Server) {
	ip := getOutBoundAddr()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:0", ip))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	port := lis.Addr().(*net.TCPAddr).Port
	ipAddrAndPort := fmt.Sprintf("%s:%d", ip, port)
	node := node.New(ipAddrAndPort)

	s := grpc.NewServer()
	pb.RegisterChordServer(s, &ChordServer{})

	// Spin up server in its own go routine to
	// avoid blocking the main routine
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	return node, s
}

//// Package private helpers

// Adapted from https://stackoverflow.com/a/37382208
func getOutBoundAddr() net.IP {
	// Get preferred outbound ip of this machine
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
