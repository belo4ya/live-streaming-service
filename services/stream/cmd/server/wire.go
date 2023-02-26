//go:build wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/stream/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/server"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
