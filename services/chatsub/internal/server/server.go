package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	v1 "github.com/belo4ya/live-streaming-service/api/chatsub/v1"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/conf"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/resolver"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"net/http"
)

var ProviderSet = wire.NewSet(NewServer)

func NewServer(c *conf.Server, r *resolver.Resolver, logger log.Logger) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	opts = append(opts, khttp.Address(c.Addr))
	srv := khttp.NewServer(opts...)

	gql := handler.NewDefaultServer(v1.NewExecutableSchema(v1.Config{Resolvers: r}))
	http.Handle("/", playground.Handler("GraphQL playground", "/chat"))
	http.Handle("/chat", gql)
	return srv
}
