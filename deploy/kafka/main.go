package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
	"time"
)

func main() {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	subscriber1, _ := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"localhost:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			//ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	subscriber2, _ := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"localhost:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			//ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)

	messages1, _ := subscriber1.Subscribe(context.Background(), "example.topic")
	messages2, _ := subscriber2.Subscribe(context.Background(), "example.topic")
	go process(messages1, 1)
	go process(messages2, 2)

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publishMessages(publisher)
}

func publishMessages(publisher message.Publisher) {
	i := 0
	for {
		payload := []byte(fmt.Sprintf("[%d] Hello, world!", i))
		msg := message.NewMessage(watermill.NewUUID(), payload)

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		i += 1
		time.Sleep(1 * time.Second)
	}
}

func process(messages <-chan *message.Message, i int) {
	for msg := range messages {
		log.Printf("[%d] received message: %s, payload: %s", i, msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
