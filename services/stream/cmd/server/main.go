package main

import (
	"github.com/belo4ya/live-streaming-service/services/stream/internal/server"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/service"
	"log"
)

func main() {
	s := service.NewService()
	srv := server.New(s, "localhost:8091")
	err := srv.Run()
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
}
