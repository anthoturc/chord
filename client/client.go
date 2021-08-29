package client

import (
	"context"
	"time"

	pb "github.com/anthoturc/chord/proto"
	"google.golang.org/grpc"
)

func CallPing(remoteAddr string) (string, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	ctx, cancel := getContext()
	defer cancel()

	resp, err := client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		return "", err
	}

	return resp.GetMessage(), nil
}

func CallFindSuccessor(remoteAddr, id string) (string, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	ctx, cancel := getContext()
	defer cancel()

	resp, err := client.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: id})
	if err != nil {
		return "", err
	}

	return resp.GetAddress(), nil
}

func CallGetPredecessor(remoteAddr string) (string, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	ctx, cancel := getContext()
	defer cancel()

	resp, err := client.GetPredecessor(ctx, &pb.GetPredecessorRequest{})
	if err != nil {
		return "", err
	}

	return resp.GetAddress(), nil
}

func CallNotify(remoteAddr, ipAddr string) error {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewChordClient(conn)

	ctx, cancel := getContext()
	defer cancel()

	_, err = client.Notify(ctx, &pb.NotifyRequest{Address: ipAddr})
	if err != nil {
		return err
	}

	return nil
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}
