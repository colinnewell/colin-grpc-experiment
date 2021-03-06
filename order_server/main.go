package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/colinnewell/grpc-experiment/orders"
)

const (
	port = ":3003"
)

type server struct {
	orders.UnimplementedStoreServer
}

func (*server) PlaceOrder(context.Context, *orders.Order) (*orders.OrderConfirmation, error) {
	return &orders.OrderConfirmation{
		OrderID: "1",
		Total:   3.33,
	}, nil
}

func (*server) PlaceOrderLines(stream orders.Store_PlaceOrderLinesServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&orders.OrderConfirmation{})
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s x %d", l.Product, l.Quantity)
	}
	// receive the lines
	// then send and close?
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	orders.RegisterStoreServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
