package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	js, _ := nc.JetStream()

	//sub, _ := js.SubscribeSync("test.subject", nats.Durable("durable-one"), nats.ManualAck())
	fmt.Println("consumer 2 listening...")

	count := 1

	js.Subscribe("t", func(msg *nats.Msg) {
		//		if count%100000 == 0 {
		fmt.Println("count", count)
		//	fmt.Println(string(msg.Data))
		//	}
		count++
		msg.Ack()

	}, nats.Durable("durablej"), nats.ManualAck())
	select {}
}
