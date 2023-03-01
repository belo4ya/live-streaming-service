// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/data"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/server"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func wireApp(http *conf.HTTP, grpc *conf.GRPC, logger log.Logger) (*kratos.App, func(), error) {
	streamServiceClient := data.NewStreamServiceClient()
	dataData, err := data.NewData(streamServiceClient)
	if err != nil {
		return nil, nil, err
	}
	streamService := service.NewStreamService(dataData, logger)
	grpcServer := server.NewGRPCServer(grpc, streamService, logger)
	clientConn := server.NewGRPCConn(grpc)
	httpServer := server.NewHTTPServer(http, clientConn, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}