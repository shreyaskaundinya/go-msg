package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/shreyaskaundinya/go-msg/pkg/controller"
	"github.com/shreyaskaundinya/go-msg/pkg/publisher"
	"github.com/shreyaskaundinya/go-msg/pkg/subscriber"
	"github.com/shreyaskaundinya/go-msg/pkg/worker"
)

func TestFirst(t *testing.T) {
	// start the controller node
	log.Println("[TEST] Creating controller")
	c := controller.CreateController("localhost", 4100)
	fmt.Println(c.Hostname)
	
	// create worker
	log.Println("[TEST] Creating worker")
	w := worker.CreateWorker("localhost", 4200)
	w.Serve()
	
	// create publisher1
	log.Println("[TEST] Creating publisher 1")
	p1 := publisher.CreatePublisher("localhost", 4201)
	fmt.Println(p1.Hostname)
	// create publisher2
	log.Println("[TEST] Creating publisher 2")
	p2 := publisher.CreatePublisher("localhost", 4202)
	fmt.Println(p2.Hostname)
	// create publisher3
	log.Println("[TEST] Creating publisher 3")
	p3 := publisher.CreatePublisher("localhost", 4203)
	fmt.Println(p3.Hostname)
	
	// create subscriber1
	log.Println("[TEST] Creating subscriber 1")
	s1:= subscriber.CreateSubscriber("localhost", 4301)
	fmt.Println(s1.Hostname)
	// create subscriber2
	log.Println("[TEST] Creating subscriber 2")
	s2 := subscriber.CreateSubscriber("localhost", 4302)
	fmt.Println(s2.Hostname)
	// create subscriber3
	log.Println("[TEST] Creating subscriber 3")
	s3 := subscriber.CreateSubscriber("localhost", 4303)
	fmt.Println(s3.Hostname)
}