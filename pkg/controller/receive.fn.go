package controller

import (
	"encoding/json"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
	"github.com/shreyaskaundinya/go-msg/pkg/queue"
	"go.uber.org/zap"
)

func (c *Controller) ReceiveMessage(msg string) error {
	// unmarshall message
	var m message.Message

	err := json.Unmarshal([]byte(msg), &m)

	if err != nil {
		zap.L().Error("[CONTROLLER] Error unmarshalling message", zap.Error(err))
		return err
	}

	// zap.L().Sugar().Infof("[CONTROLLER] Received message %v", m)

	// get topic
	t := m.Topic

	// create topic if not exists
	c.TopicsLock.Lock()
	_, ok := c.Topics[t]
	if !ok {
		zap.L().Sugar().Infof("[CONTROLLER] Topic not found: %s", t)
		c.Topics[t] = queue.NewQueue(t)
		// zap.L().Sugar().Infof("[CONTROLLER] Created topic: %v", c.Topics[t])
	}
	c.TopicsLock.Unlock()

	// enqueue message
	c.TopicsLock.RLock()
	defer c.TopicsLock.RUnlock()
	q := c.Topics[t]

	q.Send(m)

	return nil
}
