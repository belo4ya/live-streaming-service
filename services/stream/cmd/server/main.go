package main

import (
	"context"
	"github.com/belo4ya/live-streaming-service/api/stream/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	v1.UnimplementedEchoServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Echo(_ context.Context, req *v1.EchoRequest) (*v1.EchoResponse, error) {
	return &v1.EchoResponse{Value: req.Value}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	v1.RegisterEchoServiceServer(s, NewServer())
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
