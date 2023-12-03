package cmd

import (
	"github.com/shreyaskaundinya/go-msg/pkg/message"
	"github.com/shreyaskaundinya/go-msg/pkg/queue"
	"github.com/shreyaskaundinya/go-msg/pkg/utils"
	"go.uber.org/zap"
)

func TestQueue() {
	utils.InitLogger()

	q := queue.NewQueue("test")

	q.Send(message.Message{
		Topic: "test",
		Key:   "test1",
		Value: "test1",
	})

	q.Send(message.Message{
		Topic: "test",
		Key:   "test2",
		Value: "test2",
	})

	q.LogInfo()

	q.Send(message.Message{
		Topic: "test",
		Key:   "test3",
		Value: "test3",
	})

	q.LogInfo()

	zap.L().Sugar().Info(q.Read())

	q.Clear()

	q.LogInfo()
}
