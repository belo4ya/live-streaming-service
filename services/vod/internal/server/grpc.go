package server

import (
	v1 "github.com/belo4ya/live-streaming-service/api/vod/v1"
	"github.com/belo4ya/live-streaming-service/services/vod/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	grpc *grpc.Server
	addr string
}

func NewGRPCServer(s *service.Service, addr string) *GRPCServer {
	srv := grpc.NewServer()
	v1.RegisterVODServiceServer(srv, s)
	return &GRPCServer{srv, addr}
}

func (srv *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", srv.addr)
	if err != nil {
		return err
	}
	go func() {
		log.Fatalln(srv.grpc.Serve(lis))
	}()
	return nil
}
