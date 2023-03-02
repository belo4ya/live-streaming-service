package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	v1 "github.com/belo4ya/live-streaming-service/api/gqlgw/v1"
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/resolver"
	"log"
	"net/http"
	"time"
)

func main() {
	port := "8001"

	srv := handler.NewDefaultServer(v1.NewExecutableSchema(v1.Config{Resolvers: &resolver.Resolver{}}))
	srv.AddTransport(transport.Websocket{PingPongInterval: 10 * time.Second})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
