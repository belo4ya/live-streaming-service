//go:build wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/data"
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/resolver"
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, resolver.ProviderSet, newApp))
}
