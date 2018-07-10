package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	uuid "github.com/satori/go.uuid"
	pb "github.com/skjune12/gocodes/grpc/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)
	createOrders(client)
}

func createOrders(client pb.OrderServiceClient) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	order := &pb.Order{
		Id:        id.String(),
		Status:    "Created",
		CreatedOn: time.Now().Unix(),
		OrderItems: []*pb.Order_OrderItem{
			&pb.Order_OrderItem{
				Code:      "knd100",
				Name:      "Kindle Voyage",
				UnitPrice: 220,
				Quantity:  1,
			},
			&pb.Order_OrderItem{
				Code:      "kc101",
				Name:      "Kindle Voyage SmartShell Case",
				UnitPrice: 10,
				Quantity:  2,
			},
		},
	}
	res, err := client.CreateOrder(context.Background(), order)

	fmt.Printf("Result: %#v", res)
	fmt.Printf("Error: %#v", err)
}
