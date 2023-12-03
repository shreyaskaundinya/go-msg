package cmd

import (
	"github.com/shreyaskaundinya/go-msg/pkg/controller"
	"github.com/shreyaskaundinya/go-msg/pkg/utils"
)

func TestController() {
	utils.InitLogger()

	go func() {
		c := controller.NewController("localhost", 8092)

		c.Start()
	}()
}
