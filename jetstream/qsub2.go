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

	counts := 1

	js.QueueSubscribe("t", "test", func(msg *nats.Msg) {

		// if counts%100000 == 0 {
		// 	fmt.Println("count", counts)
		// 	//	fmt.Println(string(msg.Data))
		// }
		//	if counts%99000 == 0 {
		fmt.Println("cnt ", counts)
		//}
		msg.Ack()
		counts++

	}, nats.Durable("durableq"), nats.ManualAck()) //, nats.MaxAckPending(1000000))

	select {}
}
