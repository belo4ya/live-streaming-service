package server

import (
	"github.com/belo4ya/live-streaming-service/services/gateway/internal/service"
	"log"
	"net/http"
)

type Server struct {
	grpc *GRPCServer
	http *http.Server
}

func NewServer(grpcAddr string, httpAddr string, streamSvc *service.StreamService, vodSvc *service.VODService) *Server {
	grpcSrv := NewGRPCServer(grpcAddr, streamSvc, vodSvc)
	httpSrv := NewHTTPServer(httpAddr, grpcAddr)
	return &Server{grpcSrv, httpSrv}
}

func (srv *Server) RunServer() {
	log.Printf("Serving gRPC on %s\n", srv.grpc.addr)
	err := srv.grpc.Run()
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	log.Printf("Serving gRPC-Gateway on %s\n", srv.http.Addr)
	log.Fatalln(srv.http.ListenAndServe())
}
