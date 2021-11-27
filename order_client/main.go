package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/colinnewell/grpc-experiment/orders"
)

const (
	address = "localhost:3003"
	timeout = time.Duration(100 * time.Millisecond)
)

func main() {
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(timeout),
	)
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
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	r, err := client.PlaceOrder(ctx, &o)
	cancel()
	if err != nil {
		log.Fatalf("Failed to place order: %s", err)
	}
	log.Printf("%s: %f", r.GetOrderID(), r.GetTotal())
}
