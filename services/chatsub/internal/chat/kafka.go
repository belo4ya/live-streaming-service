package chat

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/conf"
)

func NewKafkaSubscriber(c *conf.Kafka) (*kafka.Subscriber, error) {
	config := kafka.DefaultSaramaSubscriberConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	return kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               c.Brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: config,
		},
		watermill.NewStdLogger(false, false),
	)
}

func NewKafkaPublisher(c *conf.Kafka) (*kafka.Publisher, error) {
	return kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   c.Brokers,
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
}
