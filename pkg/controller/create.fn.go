package controller

import (
	"github.com/shreyaskaundinya/go-msg/pkg/queue"
	"go.uber.org/zap"
)

func NewController(hostname string, port int) Controller {
	t := newServer(hostname, port)

	zap.L().Sugar().Infof("[CONTROLLER] Created controller on %s:%d", hostname, port)

	return Controller{
		Hostname: hostname,
		Port:     port,
		Topics:   make(map[string]queue.Queue),
		TCPS:     t,
	}
}

func (c *Controller) Start() {
	c.TCPS.Serve()
}
