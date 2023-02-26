package main

import (
	"github.com/belo4ya/live-streaming-service/services/vod/internal/server"
	"github.com/belo4ya/live-streaming-service/services/vod/internal/service"
)

func main() {
	s := service.NewService()
	srv := server.NewServer(s, "localhost:8092", "localhost:8082")
	srv.RunServer()
}
