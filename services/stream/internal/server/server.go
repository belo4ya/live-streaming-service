package server

import (
	v1 "github.com/belo4ya/live-streaming-service/api/stream/v1"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewServer)

func NewServer(c *conf.Server, s *service.StreamService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	opts = append(opts, grpc.Address(c.Addr))
	srv := grpc.NewServer(opts...)
	v1.RegisterStreamServiceServer(srv, s)
	return srv
}
