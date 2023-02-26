package main

import (
	"github.com/belo4ya/live-streaming-service/services/stream/internal/server"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/service"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	s := service.NewService(sugar)
	srv := server.New(s, "localhost:8091")
	err := srv.Run()
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
}
