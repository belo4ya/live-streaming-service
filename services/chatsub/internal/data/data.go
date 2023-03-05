package data

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSubscriber, NewPublisher)

func NewSubscriber() (*kafka.Subscriber, error) {
	c := kafka.DefaultSaramaSubscriberConfig()
	c.Consumer.Offsets.Initial = sarama.OffsetNewest
	return kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"localhost:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: c,
		},
		watermill.NewStdLogger(false, false),
	)
}

func NewPublisher() (*kafka.Publisher, error) {
	return kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
}
