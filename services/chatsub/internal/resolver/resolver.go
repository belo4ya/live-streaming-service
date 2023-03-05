//go:generate go run github.com/99designs/gqlgen generate

package resolver

import (
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewResolver)

type Resolver struct {
	chats ChatMap
	s     *kafka.Subscriber
	p     *kafka.Publisher
	log   *log.Helper
}

func NewResolver(s *kafka.Subscriber, p *kafka.Publisher, logger log.Logger) *Resolver {
	return &Resolver{
		chats: NewChatMap(),
		s:     s,
		p:     p,
		log:   log.NewHelper(logger),
	}
}
