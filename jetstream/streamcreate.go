package main

import (

	//	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(100))

	js.AddStream(&nats.StreamConfig{
		Name:     "TEST",
		Subjects: []string{"t"},
		//MaxMsgs:   100000,
		Storage:  nats.MemoryStorage,
		Replicas: 1,
		MaxAge:   3600 * time.Second,
		//Retention: nats.WorkQueuePolicy,
	})

}
