package server

import (
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	"github.com/belo4ya/live-streaming-service/services/gateway/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	grpc *grpc.Server
	addr string
}

func NewGRPCServer(addr string, streamSvc *service.StreamService, vodSvc *service.VODService) *GRPCServer {
	srv := grpc.NewServer()
	v1.RegisterStreamServiceServer(srv, streamSvc)
	v1.RegisterVODServiceServer(srv, vodSvc)
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
