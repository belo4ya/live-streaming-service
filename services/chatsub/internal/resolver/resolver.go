//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/chat"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewResolver)

type Resolver struct {
	chat *chat.Controller
	log  *log.Helper
}

func NewResolver(c *chat.Controller, logger log.Logger) *Resolver {
	return &Resolver{chat: c, log: log.NewHelper(logger)}
}
