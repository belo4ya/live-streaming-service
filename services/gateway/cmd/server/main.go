package main

import (
	"github.com/belo4ya/live-streaming-service/services/gateway/internal/server"
	"github.com/belo4ya/live-streaming-service/services/gateway/internal/service"
)

func main() {
	streamSvc := service.NewStreamService()
	vodSvc := service.NewVODService()
	srv := server.NewServer("localhost:8090", "localhost:8080", streamSvc, vodSvc)
	srv.RunServer()
}
