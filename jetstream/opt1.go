package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	// Create JetStream context
	js, _ := nc.JetStream()

	// Create the stream if not already created
	_, err := js.AddStream(&nats.StreamConfig{
		Name:      "TEST",
		Subjects:  []string{"fas"},
		MaxMsgs:   5000,
		Storage:   nats.FileStorage, // Persistent storage
		Replicas:  1,
		MaxAge:    3600 * time.Second, // Retain for 1 hour
		Retention: nats.LimitsPolicy,  // Retain messages until limits are hit
		Discard:   nats.DiscardOld,    // Discard old messages if the stream is full
	})

	if err != nil {
		fmt.Println("Error creating stream:", err)
		return
	}

	// Sample JSON message to publish
	s := []byte(`
		"nfInstanceId": "a52f47c0upf",
		"nfType": "UDM",
		"nfStatus": "REGISTERED",
		"heartBeatTimer": 20
	`)

	// Publish 10 messages asynchronously
	for i := 0; i < 10; i++ {
		msg := &nats.Msg{
			Subject: "fas",
			Data:    s,
		}
		js.PublishMsgAsync(msg)
		time.Sleep(1 * time.Second) // Simulating delay
	}

	// Ensure all messages have been published
	js.PublishAsyncComplete()

	nc.Flush()
}
