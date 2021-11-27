package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/colinnewell/grpc-experiment/orders"
)

const (
	port = ":3003"
)

type server struct {
	orders.UnimplementedOrderServerServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	orders.RegisterOrderServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
