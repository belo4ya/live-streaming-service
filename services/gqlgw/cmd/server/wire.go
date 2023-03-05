//go:build wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/data"
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/resolver"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireResolver(log.Logger) (*resolver.Resolver, func(), error) {
	panic(wire.Build(data.ProviderSet, resolver.NewResolver))
}
