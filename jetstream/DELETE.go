package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer nc.Drain()

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("Error getting JetStream context: %v", err)
	}

	// Delete the stream
	streamName := "TEST"
	err = js.DeleteStream(streamName)
	if err != nil {
		log.Fatalf("Error deleting stream %q: %v", streamName, err)
	}

	log.Printf("Stream %q deleted successfully", streamName)
}
