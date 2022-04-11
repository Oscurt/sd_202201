package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	pb      "app/config"
)

var port = flag.Int("port", 50051, "The server port")
type server struct {}

func (s *server) GetItem(ctx context.Context, in *pb.GetItemsRequest) (*pb.Response, error) {
	log.Printf("Received: %v", in)
}

func GrpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterItemServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}