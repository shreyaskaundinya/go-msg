package cmd

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/shreyaskaundinya/go-msg/pkg/controller"
	"github.com/shreyaskaundinya/go-msg/pkg/publisher"
	"github.com/shreyaskaundinya/go-msg/pkg/utils"
)

func TestController() {
	f, err := os.OpenFile("S:/Documents/PROJECTS/go-msg/random/app.log", os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	utils.InitLogger(f)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		c := controller.NewController("localhost", 8092)

		c.Start()
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {

			p := publisher.NewPublisher(
				[]string{"localhost:8092"},
			)

			// zap.L().Sugar().Info("Sending message to controller")

			for j := 0; j < 20000; j++ {
				p.SendKV("test", fmt.Sprintf("%d", i), fmt.Sprintf("%d", j))
			}

			p.Close()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
