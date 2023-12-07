package publisher

import (
	"net"
	"time"

	"go.uber.org/zap"
)

func (p *Publisher) connect() {
	conn, err := net.Dial("tcp", p.controllerServers[0])

	if err != nil {
		panic(err)
	}

	p.conn = conn
}

func (p *Publisher) Close() {
	// wait for exit channel
	// p.bufferClearChan <- true

	// for t := range p.exitChan {
	// 	if t {
	// 		break
	// 	}
	// }

	// close(p.exitChan)
	// close(p.bufferClearChan)
	// p.wg.Wait()
	close(p.sendBufferChan)

	p.wg.Wait()

	time.Sleep(5 * time.Second)
	p.conn.Close()
	zap.L().Info("[PUBLISHER] Publisher closed")
}
