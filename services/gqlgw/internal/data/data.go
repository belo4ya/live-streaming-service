package data

import (
	"context"
	streamv1 "github.com/belo4ya/live-streaming-service/api/stream/v1"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewStreamServiceClient)

type Data struct {
	StreamC streamv1.StreamServiceClient
}

func NewData(sc streamv1.StreamServiceClient) (*Data, error) {
	return &Data{StreamC: sc}, nil
}

func NewStreamServiceClient() streamv1.StreamServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9001"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return streamv1.NewStreamServiceClient(conn)
}
