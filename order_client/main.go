package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/colinnewell/grpc-experiment/orders"
)

const address = "localhost:3003"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := orders.NewOrderServerClient(conn)
	o := orders.Order{
		OrderRef: "MyOrder1",
		Address:  "London",
		Notes:    "Deliver to car park",
	}
	r, err := client.PlaceOrder(context.TODO(), &o)
	if err != nil {
		log.Fatalf("Failed to place order: %s", err)
	}
	log.Printf("%s: %f", r.GetOrderID(), r.GetTotal())
}
