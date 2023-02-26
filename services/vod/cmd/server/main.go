package main

import (
	"github.com/belo4ya/live-streaming-service/services/vod/internal/server"
	"github.com/belo4ya/live-streaming-service/services/vod/internal/service"
	"log"
)

func main() {
	s := service.NewService()
	srv := server.New(s, "localhost:8092")
	err := srv.Run()
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
}
