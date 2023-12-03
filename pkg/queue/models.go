package queue

import (
	"sync"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
)

type Queue struct {
	Data        []message.Message
	Topic       string
	Top         int
	Lock        sync.RWMutex
	EnqueueChan chan message.Message
}
