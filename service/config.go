package service

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/anthoturc/chord/node"
	"google.golang.org/grpc"

	pb "github.com/anthoturc/chord/proto"
)

type ChordServerConfig struct {
	Create           bool
	Join             bool
	RemoteNodeIpAddr string
}

//// Chord service initialization
func Init(config *ChordServerConfig) *ChordServer {
	config.validate()

	ip := getOutBoundAddr()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:0", ip))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	port := lis.Addr().(*net.TCPAddr).Port
	ipAddrAndPort := fmt.Sprintf("%s:%d", ip, port)
	node := node.New(ipAddrAndPort)

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
	chordServer := &ChordServer{node: node, server: s}
	pb.RegisterChordServer(s, chordServer)
	log.Printf("Staring up ChordServer on: %s\n", ipAddrAndPort)
	startGrpcServer(s, lis, ipAddrAndPort)
	startStabilizingRoutines(node)

	return chordServer
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

	go func() {
		for {
			node.FixFingers()
			time.Sleep(2 * time.Second)
		}
	}()
}

func (c *ChordServerConfig) validate() {
	if c.Create == c.Join {
		log.Fatal("Cannot create a chord ring and join one too")
	}
}
