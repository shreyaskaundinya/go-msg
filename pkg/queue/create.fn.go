package queue

import (
	"sync"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
	"go.uber.org/zap"
)

func NewQueue(topic string) *Queue {
	q := &Queue{
		Top:         -1,
		Data:        make([]message.Message, 0),
		Topic:       topic,
		Lock:        sync.RWMutex{},
		EnqueueChan: make(chan message.Message),
	}

	zap.L().Sugar().Infof("[QUEUE] Created queue %s", topic)
	zap.L().Sugar().Infof("[QUEUE] Starting enqueue routine for %s", topic)

	go func() {
		for msg := range q.EnqueueChan {
			q.enqueue(msg)
		}
	}()

	return q
}
