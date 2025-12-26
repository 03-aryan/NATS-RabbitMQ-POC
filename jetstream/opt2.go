// package main

// import (
// 	"fmt"

// 	"github.com/nats-io/nats.go"
// )

// func main() {
// 	// Connect to NATS server
// 	nc, _ := nats.Connect(nats.DefaultURL)
// 	defer nc.Drain()

// 	// Create JetStream context
// 	js, _ := nc.JetStream()

// 	// Subscribe with durable subscription and manual ack
// 	fmt.Println("Consumer listening...")

// 	count := 1
// 	_, err := js.Subscribe("fas", func(msg *nats.Msg) {
// 		// Process the message
// 		//	fmt.Println("Processing message:", string(msg.Data))
// 		fmt.Println("Message count:", count)

// 		// Acknowledge the message after processing
// 		err := msg.Ack()
// 		if err != nil {
// 			fmt.Println("Failed to acknowledge message:", err)
// 		}

// 		count++
// 	}, nats.Durable("Durable"), nats.ManualAck())

// 	// Check for errors when subscribing
// 	if err != nil {
// 		fmt.Println("Error subscribing:", err)
// 		return
// 	}

// 	// Keep the subscriber running indefinitely
// 	select {}
// }

package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	// Create JetStream context
	js, _ := nc.JetStream()

	// Print stream info to check how many messages are in the stream
	// streamInfo, err := js.StreamInfo("TEST")
	// if err != nil {
	// 	fmt.Println("Error fetching stream info:", err)
	// } else {
	// 	//fmt.Printf("Stream Info: %+v\n", streamInfo)
	// }

	// Print subscription info to check the position of the durable subscription
	// subInfo, err := js.DurableInfo("Durable")
	// if err != nil {
	// 	fmt.Println("Error fetching durable subscription info:", err)
	// } else {
	// 	fmt.Printf("Durable Subscription Info: %+v\n", subInfo)
	// }

	// Subscribe with durable subscription and manual ack
	fmt.Println("Consumer listening...")

	count := 1
	_, err := js.Subscribe("fas", func(msg *nats.Msg) {
		// Process the message
		//	fmt.Println("Processing message:", string(msg.Data))
		fmt.Println("Message count:", count)

		// Acknowledge the message after processing
		err := msg.Ack()
		if err != nil {
			fmt.Println("Failed to acknowledge message:", err)
		} else {
			//fmt.Println("Message acknowledged:", string(msg.Data))
			fmt.Println("Message acknowledged:")
		}

		count++
	}, nats.Durable("Durable"), nats.ManualAck())

	// Check for errors when subscribing
	if err != nil {
		fmt.Println("Error subscribing:", err)
		return
	}

	// Keep the subscriber running indefinitely
	select {}
}
