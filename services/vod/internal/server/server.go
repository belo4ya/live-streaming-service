package server

import (
	v1 "github.com/belo4ya/live-streaming-service/api/vod/v1"
	"github.com/belo4ya/live-streaming-service/services/vod/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpc *grpc.Server
	addr string
}

func New(s *service.Service, addr string) *Server {
	srv := grpc.NewServer()
	v1.RegisterVODServiceServer(srv, s)
	return &Server{srv, addr}
}

func (srv *Server) Run() error {
	lis, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}
	log.Printf("Serving gRPC on %s\n", srv.addr)
	return srv.grpc.Serve(lis)
}
