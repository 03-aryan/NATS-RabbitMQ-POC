package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	// Create a channel
	ch, _ := conn.Channel()
	defer ch.Close()

	// Declare a queue
	q, _ := ch.QueueDeclare(
		"taskqueue", // Name of the queue
		true,        // Durable (survive server restarts)
		false,       // Delete when unused
		false,       // Exclusive
		false,       // No-wait
		nil,         // Arguments
	)

	// Receive messages
	msgs, _ := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer name
		false,  // Auto acknowledge (set false for manual ack)
		false,  // Exclusive
		false,  // No local
		false,  // No wait
		nil,    // Arguments
	)

	// Start consuming
	fmt.Println(" Waiting for messages sub 1.")

	count := 0
	for msg := range msgs {
		// Print the message and count
		//	fmt.Println(" [x] Received %s \n count %d", string(msg.Body), count)
		if count%1000000 == 0 {

			fmt.Println(count)
		}
		count++
		// Acknowledge the message (manual ack)
		msg.Ack(true)
	}
}
