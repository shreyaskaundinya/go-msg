package publisher

import (
	"errors"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
	"go.uber.org/zap"
)

func (p *Publisher) flushBufferRoutine() {
	for msg := range p.sendBufferChan {
		// select {
		// case <-p.bufferClearChan:
		// 	p.exitChan <- true
		// 	break loop
		// default:
		// }
		p.wg.Add(1)
		zap.L().Sugar().Info("[PUBLISHER] Sending message to controller")
		err := p.sendToController(msg)

		if err != nil {
			// write back; retry
			zap.L().Sugar().Info("[PUBLISHER] Error sending message to controller. Retrying")
			p.sendBufferChan <- msg
		}
	}
}

func (p *Publisher) sendToController(msg message.Message) error {
	defer p.wg.Done()

	// zap.L().Sugar().Info("[PUBLISHER] Marshalling message")
	m, err := msg.Marshall()

	if err != nil {
		zap.L().Error("[PUBLISHER] Error marshalling message", zap.Error(err))
		return err
	}

	// zap.L().Sugar().Info("[PUBLISHER] Appending \\n")
	m = append(m, '\n')

	_, err = p.conn.Write(m)

	if err != nil {
		zap.L().Error("[PUBLISHER] Error sending message to controller", zap.Error(err))
		return err
	}
	zap.L().Sugar().Info("[PUBLISHER] Wrote to controller")

	// zap.L().Sugar().Info("[PUBLISHER] Reading response from controller")

	buf := make([]byte, 1024)
	n, err := p.conn.Read(buf)

	if err != nil {
		zap.L().Error("[PUBLISHER] Error reading response from controller", zap.Error(err))
		return err
	}

	// zap.L().Sugar().Info("[PUBLISHER] Response from controller", string(buf[:n]), n)

	if string(buf[:n]) != "ACK\n" {
		// zap.L().Error("[PUBLISHER] Error reading response from controller", zap.Error(err))
		return errors.New("error reading response from controller")
	}

	return nil
}

func (p *Publisher) SendKV(topic string, key string, value string) {
	msg := message.Message{
		Topic:  topic,
		Key:    key,
		Value:  value,
		Offset: -1,
	}

	p.sendBufferChan <- msg
}

func (p *Publisher) Send(topic string, value string) {
	p.SendKV(topic, "", value)
}
