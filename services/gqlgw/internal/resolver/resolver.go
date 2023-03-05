//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"github.com/belo4ya/live-streaming-service/services/gqlgw/internal/data"
	"github.com/go-kratos/kratos/v2/log"
)

type Resolver struct {
	data *data.Data
	log  *log.Helper
}

func NewResolver(data *data.Data, logger log.Logger) *Resolver {
	return &Resolver{data: data, log: log.NewHelper(logger)}
}
