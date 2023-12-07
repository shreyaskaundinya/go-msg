package controller

import (
	"sync"

	"github.com/shreyaskaundinya/go-msg/pkg/queue"
	"go.uber.org/zap"
)

func NewController(hostname string, port int) Controller {
	t := newServer(hostname, port)

	zap.L().Sugar().Infof("[CONTROLLER] Created controller on %s:%d", hostname, port)

	return Controller{
		Hostname:   hostname,
		Port:       port,
		Topics:     make(map[string]*queue.Queue),
		TCPS:       t,
		TopicsLock: sync.RWMutex{},
	}
}

func (c *Controller) Start() {
	c.Serve()
}
