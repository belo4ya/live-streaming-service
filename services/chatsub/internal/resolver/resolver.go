//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewResolver)

type Resolver struct {
	log *log.Helper
}

func NewResolver(logger log.Logger) *Resolver {
	return &Resolver{log: log.NewHelper(logger)}
}
