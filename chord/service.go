package chord

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/anthoturc/chord/hash"
	"github.com/anthoturc/chord/node"

	pb "github.com/anthoturc/chord/proto"
	"google.golang.org/grpc"
)

type ChordServerConfig struct {
	Create           bool
	Join             bool
	RemoteNodeIpAddr string
}

type ChordServer struct {
	// TODO: Add fields that the server will need here
	// e.g this could be data base connections
	Node *node.ChordNode

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
	localNode := s.Node
	successorAddr, err := localNode.FindSuccessor(hash.Hash(req.GetId()))
	return &pb.FindSuccessorResponse{Address: successorAddr}, err
}

func (s *ChordServer) GetPredecessor(ctx context.Context, req *pb.GetPredecessorRequest) (*pb.GetPredecessorResponse, error) {
	localNode := s.Node
	if localNode.Predecessor == nil {
		return nil, errors.New("no predecessor for this node")
	}

	return &pb.GetPredecessorResponse{Address: localNode.Predecessor.IpAddr}, nil
}

func (s *ChordServer) Notify(ctx context.Context, req *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	localNode := s.Node
	localNode.Notify(req.GetAddress())

	return &pb.NotifyResponse{}, nil
}

////

//// Chord service initialization
func Init(config *ChordServerConfig) (*node.ChordNode, *grpc.Server) {
	ip := getOutBoundAddr()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:0", ip))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	port := lis.Addr().(*net.TCPAddr).Port
	ipAddrAndPort := fmt.Sprintf("%s:%d", ip, port)
	node := node.New(ipAddrAndPort)

	config.validate()
	if config.Create { // on create set the successor to itself.
		log.Println("Creating new ChordRing...")
		node.Successors[1] = node
	} else { // on join, make RPC find predecessor call
		log.Println("Attempting to join ChordRing...")
		err := node.Join(config.RemoteNodeIpAddr)
		if err != nil {
			log.Fatalf("Error joining chord ring: %v", err)
		}
	}

	s := grpc.NewServer()
	pb.RegisterChordServer(s, &ChordServer{Node: node})
	log.Printf("Staring up ChordServer on: %s\n", ipAddrAndPort)
	startGrpcServer(s, lis, ipAddrAndPort)
	startStabilizingRoutines(node)

	return node, s
}

//// Package private helpers

func startGrpcServer(s *grpc.Server, lis net.Listener, ipAddrAndPort string) {
	// Spin up server in its own go routine to
	// avoid blocking the main routine
	go func() {
		log.Fatal(s.Serve(lis))
	}()
}

func startStabilizingRoutines(node *node.ChordNode) {
	go func() {
		for {
			node.CheckPredecessor()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			node.Stabilize()
			time.Sleep(2 * time.Second)
		}
	}()
}

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

func (c *ChordServerConfig) validate() {
	if c.Create == c.Join {
		log.Fatal("Cannot create a chord ring and join one too")
	}
}
