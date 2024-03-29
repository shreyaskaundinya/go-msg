package controller

import (
	"sync"

	"github.com/shreyaskaundinya/go-msg/pkg/queue"
)

type Controller struct {
	// hostname
	Hostname string
	// port
	Port int

	// topics
	Topics map[string]*queue.Queue

	// lock
	TopicsLock sync.RWMutex

	// TCP
	TCPS *TCPServer
}
