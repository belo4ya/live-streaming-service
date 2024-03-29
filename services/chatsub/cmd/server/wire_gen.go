// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/chat"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/resolver"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/server"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func wireChatController(kafka *conf.Kafka, logger log.Logger) (*chat.Broadcaster, func(), error) {
	subscriber, err := chat.NewKafkaSubscriber(kafka)
	if err != nil {
		return nil, nil, err
	}
	publisher, err := chat.NewKafkaPublisher(kafka)
	if err != nil {
		return nil, nil, err
	}
	broadcaster := chat.NewBroadcaster(subscriber, publisher, kafka, logger)
	return broadcaster, func() {
	}, nil
}

func wireApp(confServer *conf.Server, broadcaster *chat.Broadcaster, logger log.Logger) (*kratos.App, func(), error) {
	resolverResolver := resolver.NewResolver(broadcaster, logger)
	httpServer := server.NewServer(confServer, resolverResolver, logger)
	app := newApp(logger, httpServer)
	return app, func() {
	}, nil
}
