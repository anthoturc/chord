package client

import (
	"context"

	pb "github.com/anthoturc/chord/proto"
	"google.golang.org/grpc"
)

func CallPing(remoteAddr string) (string, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	resp, err := client.Ping(context.Background(), &pb.PingRequest{})
	if err != nil {
		return "", err
	}

	return resp.GetMessage(), nil
}

func CallFindSuccessor(remoteAddr, id string) (string, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	resp, err := client.FindSuccessor(context.Background(), &pb.FindSuccessorRequest{Id: id})
	if err != nil {
		return "", err
	}

	return resp.GetAddress(), nil
}

func CallGetPredecessor(remoteAddr string) (string, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	resp, err := client.GetPredecessor(context.Background(), &pb.GetPredecessorRequest{})
	if err != nil {
		return "", err
	}

	return resp.GetAddress(), nil
}

func CallNotify(remoteAddr, ipAddr string) error {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	_, err = client.Notify(context.Background(), &pb.NotifyRequest{Address: ipAddr})
	if err != nil {
		return err
	}

	return nil
}
