package publisher

import (
	"net"
	"sync"

	"github.com/shreyaskaundinya/go-msg/pkg/message"
)

type Publisher struct {
	// controller
	controllerServers []string

	// buffer of messages to be sent
	sendBufferChan chan message.Message

	// wg
	wg sync.WaitGroup

	// conn
	conn net.Conn
}
