package cmd

import (
	"os"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
	"github.com/shreyaskaundinya/go-msg/pkg/queue"
	"github.com/shreyaskaundinya/go-msg/pkg/utils"
	"go.uber.org/zap"
)

func TestQueue() {
	f, err := os.OpenFile("S:/Documents/PROJECTS/go-msg/random/app.log", os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	utils.InitLogger(f)

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
