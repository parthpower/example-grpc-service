// Package servicea implements a server for Greeter service.
package servicea

import (
	"context"
	"log"
	"net"

	pb "examplesvc/pkg/api/serviceav1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedEchoAServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	log.Printf("Received: %v", in.GetMsg())
	return &pb.PingResponse{}, nil
}

// Serve start servicea server
func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoAServer(s, &server{})
	reflection.Register(s)
	log.Printf("Listening")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
