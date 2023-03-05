package main

import (
	"flag"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"os"
)

// go build -ldflags "-X main.Version=0.0.1 -X main.Name=stream"
var (
	Version = "0.0.1"
	Name    = "stream"
	id, _   = os.Hostname()
)

func newApp(logger log.Logger, srv *grpc.Server) *kratos.App {
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
		"services/stream/configs/config.yaml",
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

	app, cleanup, err := wireApp(b.Server, b.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
