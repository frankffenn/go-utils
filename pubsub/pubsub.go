package sectorstorage

import (
	"context"

	"github.com/whyrusleeping/pubsub"
)

type PubSub struct {
	incoming *pubsub.PubSub
}

func NewPubSub(capacity int) *PubSub {
	return &PubSub{incoming: pubsub.New(capacity)}
}

func (s *PubSub) Subscribe(ctx context.Context, topics ...string) (<-chan interface{}, error) {
	sub := s.incoming.Sub(topics...)
	out := make(chan interface{}, 10)

	go func() {
		defer s.incoming.Unsub(sub, topics...)

		for {
			select {
			case r := <-sub:
				select {
				case out <- r:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, nil
}

func (s *PubSub) Publish(ctx context.Context, msg interface{}, topics ...string) {
	s.incoming.Pub(msg, topics...)
}
