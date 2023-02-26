package main

import (
	"context"
	streamv1 "github.com/belo4ya/live-streaming-service/api/stream/v1"
	vod1 "github.com/belo4ya/live-streaming-service/api/vod/v1"
	"github.com/belo4ya/live-streaming-service/services/gateway/internal/server"
	"github.com/belo4ya/live-streaming-service/services/gateway/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8091",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	streamSvc := service.NewStreamService(streamv1.NewStreamServiceClient(conn))

	conn, err = grpc.DialContext(
		context.Background(),
		"localhost:8092",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	vodSvc := service.NewVODService(vod1.NewVODServiceClient(conn))

	srv := server.NewServer("localhost:8090", "localhost:8080", streamSvc, vodSvc)
	srv.Run()
}
