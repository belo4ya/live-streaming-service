package main

import (
	"github.com/belo4ya/live-streaming-service/services/stream/internal/server"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/service"
)

func main() {
	s := service.NewService()
	srv := server.NewServer(s, "localhost:8091", "localhost:8081")
	srv.RunServer()
}
