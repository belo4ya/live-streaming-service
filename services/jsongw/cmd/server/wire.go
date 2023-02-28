//go:build wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/data"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/server"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.HTTP, *conf.GRPC, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
