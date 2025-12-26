package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("error in connect %v", err)
	}
	defer nc.Drain()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}

	fmt.Println("consumer 1 listening...")

	counts := 1
	//var mu sync.Mutex
	js.Subscribe("t", func(msg *nats.Msg) {
		//	mu.Lock()
		//defer mu.Unlock()

		//if counts%1000 == 0 {
		fmt.Println("count", counts, msg.Header.Get("Head"))
		//}
		trackingmsg := "sub1"
		nc.Publish("tracking.processed", []byte(trackingmsg))
		msg.Ack()
		counts++

	}, nats.Durable("durable_first"), nats.ManualAck(), nats.DeliverAll())
	select {}
}
