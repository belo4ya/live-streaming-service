package server

import (
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *conf.GRPC, s *service.StreamService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	opts = append(opts, grpc.Address(c.Addr))
	srv := grpc.NewServer(opts...)
	v1.RegisterStreamServiceServer(srv, s)
	return srv
}
