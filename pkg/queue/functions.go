package queue

import (
	"github.com/shreyaskaundinya/go-msg/pkg/message"
	"go.uber.org/zap"
)

func (q *Queue) Send(msg message.Message) {
	q.EnqueueChan <- msg
}

func (q *Queue) enqueue(item message.Message) {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	zap.L().Sugar().Infof("[QUEUE] Enqueueing message %v", item)

	q.Data = append(q.Data, item)
	q.Top++
}

func (q *Queue) Read() []message.Message {
	q.Lock.RLock()
	defer q.Lock.RUnlock()
	return q.Data[:q.Top+1]
}

func (q *Queue) ReadFromOffset(offset int) []message.Message {
	q.Lock.RLock()
	defer q.Lock.RUnlock()
	return q.Data[offset:]
}

func (q *Queue) Clear() {
	q.Lock.Lock()
	defer q.Lock.Unlock()
	q.Data = nil
	q.Top = -1
}

func (q *Queue) LogInfo() {
	zap.L().Sugar().Infof("[QUEUE] %s : has %d messages", q.Topic, q.Top+1)
}
