package main

import (
	"context"
	"flag"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"
)

// go build -ldflags "-X main.Version=0.0.1 -X main.Name=stream"
var (
	Version = "0.0.1"
	Name    = "chat-sub"
	id, _   = os.Hostname()
)

func newApp(logger log.Logger, srv *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(srv),
	)
}

func main() {
	var confPath string
	flag.StringVar(
		&confPath,
		"conf",
		"services/chatsub/configs/config.yaml",
		"config path, eg: --conf config.yaml",
	)
	flag.Parse()

	c, b, err := conf.Load(confPath)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"service", Name+":"+Version,
		"caller", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(logger)

	chat, cleanup, err := wireChatController(b.Kafka, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	app, cleanup, err := wireApp(b.Server, chat, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := chat.RunBroadcast(context.Background()); err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
