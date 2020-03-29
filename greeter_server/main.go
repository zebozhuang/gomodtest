package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/zebozhuang/gomodtest/helloworld"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}


func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}