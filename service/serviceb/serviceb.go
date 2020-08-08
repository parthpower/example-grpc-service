// Package serviceb implements a server for Greeter service.
package serviceb

import (
	"context"
	"log"
	"net"

	pb "examplesvc/pkg/api/servicebv1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedEchoBServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Flip(ctx context.Context, in *pb.FlipRequest) (*pb.FlipResponse, error) {
	log.Println("Received:", in.String())
	return &pb.FlipResponse{
		Msg:   "pong",
		Rcode: 0,
	}, nil
}

// Serve start servicea server
func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoBServer(s, &server{})
	reflection.Register(s)
	log.Printf("Listening")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
