package publisher

import (
	"sync"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
)

func NewPublisher(ControllerServers []string) *Publisher {
	p := &Publisher{
		controllerServers: ControllerServers,
		sendBufferChan:    make(chan message.Message),
		// bufferClearChan:   make(chan bool),
		// exitChan:          make(chan bool),
		wg: sync.WaitGroup{},
	}

	p.connect()
	go p.flushBufferRoutine()
	// go p.flushBufferRoutine()
	// go p.flushBufferRoutine()

	return p
}
