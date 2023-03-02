package chat

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	v1 "github.com/belo4ya/live-streaming-service/api/chatsub/v1"
	"github.com/belo4ya/live-streaming-service/services/chatsub/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"sync"
)

var ProviderSet = wire.NewSet(NewController, NewKafkaSubscriber, NewKafkaPublisher)

type Subscriber struct {
	ch chan *v1.Message
}

type Chat struct {
	subs sync.Map
}

type Controller struct {
	chats sync.Map
	s     *kafka.Subscriber
	p     *kafka.Publisher
	topic string
	log   *log.Helper
}

func NewController(s *kafka.Subscriber, p *kafka.Publisher, c *conf.Kafka, logger log.Logger) *Controller {
	return &Controller{
		chats: sync.Map{},
		s:     s,
		p:     p,
		topic: c.Topic,
		log:   log.NewHelper(logger),
	}
}

func (c *Controller) LoadOrStore(id string) (*Chat, bool) {
	chat, loaded := c.chats.LoadOrStore(id, &Chat{subs: sync.Map{}})
	return chat.(*Chat), loaded
}

func (c *Controller) Publish(msg *v1.Message) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	event := message.NewMessage(watermill.NewUUID(), payload)
	if err := c.p.Publish(c.topic, event); err != nil {
		return err
	}
	return nil
}

func (c *Controller) RunBroadcast(ctx context.Context) error {
	events, err := c.s.Subscribe(ctx, c.topic)
	if err != nil {
		return err
	}

	go func() {
		for event := range events {
			var msg v1.Message
			if err := json.Unmarshal(event.Payload, &msg); err != nil {
				c.log.Errorw(
					log.DefaultMessageKey, "Unmarshal error",
					"err", err,
					"event_id", event.UUID,
				)
				continue
			}

			chat, _ := c.LoadOrStore(msg.ChannelID)
			chat.PublishAll(&msg)
			event.Ack()
		}
	}()
	return nil
}

func (c *Chat) Store(id string, s *Subscriber) {
	c.subs.Store(id, s)
}

func (c *Chat) Delete(id string) {
	c.subs.Delete(id)
}

func (c *Chat) Range(f func(key string, value *Subscriber) bool) {
	c.subs.Range(func(k, v any) bool {
		return f(k.(string), v.(*Subscriber))
	})
}

func (c *Chat) PublishAll(msg *v1.Message) {
	c.subs.Range(func(_, v any) bool {
		sub := v.(*Subscriber)
		sub.ch <- msg
		return true
	})
}

func NewSubscriber(ch chan *v1.Message) *Subscriber {
	return &Subscriber{ch: ch}
}
