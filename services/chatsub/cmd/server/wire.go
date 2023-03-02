//go:build wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/chat"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/resolver"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireChatController(*conf.Kafka, log.Logger) (*chat.Controller, func(), error) {
	panic(wire.Build(chat.ProviderSet))
}

func wireApp(*conf.Server, *chat.Controller, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, resolver.ProviderSet, newApp))
}
