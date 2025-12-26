package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	//const n = "nats://127.0.0.1:4222"
	nc, err := nats.Connect(nats.DefaultURL) //nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS:", err)
		return
	}
	defer nc.Close()
	count := 0
	// Subscribe to subject
	_, err = nc.Subscribe("test.header", func(msg *nats.Msg) {
		//fmt.Println("Header publisher:", msg.Header.Get("1"), "Data:", string(msg.Data), "\n count: ", count)
		if count%1000000 == 0 {
			fmt.Println(count)
		}
		count++
	})
	if err != nil {
		fmt.Println("Error subscribing:", err)
		return
	}

	//nc.Flush() // Ensure subscription is ready

	fmt.Println("Waiting for messages...")

	select {}
}
