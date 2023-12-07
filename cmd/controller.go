package cmd

import (
	"fmt"
	"math/rand"
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

			for j := 0; j < 10; j++ {
				p.SendKV(fmt.Sprintf("test%d", i), fmt.Sprintf("%d", i), fmt.Sprintf("%d", j))
				// t :=
				// zap.L().Sugar().Infof("Sleeping for %v", t)
				time.Sleep(time.Duration(+rand.Float64()*50) * time.Millisecond)
			}

			p.Close()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
