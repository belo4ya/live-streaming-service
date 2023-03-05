package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	v1 "github.com/belo4ya/live-streaming-service/api/gqlgw/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"net/http"
	"os"
	"time"
)

// go build -ldflags "-X main.Version=0.0.1 -X main.Name=stream"
var (
	Version = "0.0.1"
	Name    = "json-gateway"
	id, _   = os.Hostname()
)

func main() {
	port := "8001"

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"service", Name+":"+Version,
		"caller", log.DefaultCaller,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	log.SetLogger(logger)

	r, cleanup, err := wireResolver(logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	srv := handler.NewDefaultServer(v1.NewExecutableSchema(v1.Config{Resolvers: r}))
	srv.AddTransport(transport.Websocket{PingPongInterval: 10 * time.Second})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
