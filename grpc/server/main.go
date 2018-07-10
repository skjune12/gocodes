package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"

	pb "github.com/skjune12/gocodes/grpc/pb"
)

type server struct{}

func (s *server) CreateOrder(ctx context.Context, in *pb.Order) (*pb.OrderResponse, error) {
	spew.Dump(in)
	return nil, errors.New("Not Found")
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()

	pb.RegisterOrderServiceServer(s, &server{})
	s.Serve(listen)
}
