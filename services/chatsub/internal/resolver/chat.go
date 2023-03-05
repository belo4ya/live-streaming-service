package resolver

import (
	v1 "github.com/belo4ya/live-streaming-service/api/chatsub/v1"
	"sync"
)

type Subscriber struct {
	ch chan *v1.Message
}

type Chat struct {
	subs sync.Map
}

type ChatMap struct {
	chats sync.Map
}

func (c *Chat) Add(id string, s *Subscriber) {
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

func NewChatMap() ChatMap {
	return ChatMap{chats: sync.Map{}}
}

func (m *ChatMap) GetOrCreate(id string) *Chat {
	chat, _ := m.chats.LoadOrStore(id, &Chat{subs: sync.Map{}})
	return chat.(*Chat)
}
