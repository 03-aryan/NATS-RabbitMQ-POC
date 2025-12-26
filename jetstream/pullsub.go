package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	sub, err := js.PullSubscribe("t", "durablepull_1", nats.ManualAck())
	if err != nil {
		panic(err)
	}

	count := 0
	start := time.Now()

	for count < 100_000 {
		msgs, err := sub.Fetch(1000, nats.MaxWait(20*time.Millisecond))
		if err != nil && err != nats.ErrTimeout {
			continue
		}
		for _, msg := range msgs {
			msg.Ack()
			count++
		}
	}

	fmt.Printf(" Total received: %d in %v\n", count, time.Since(start))
}
