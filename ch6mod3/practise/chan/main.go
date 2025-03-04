package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare a queue (must match the producer's declaration)
	queue, err := ch.QueueDeclare(
		"task_queue", // Name of the queue
		true,         // Durable
		false,        // Delete when unused
		false,        // Exclusive
		false,        // No-wait
		nil,          // Arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Consume messages from the queue
	msgs, err := ch.Consume(
		queue.Name, // Queue name
		"",         // Consumer tag
		true,       // Auto-acknowledge
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Arguments
	)
	failOnError(err, "Failed to register a consumer")

	// Create a channel to signal when done
	forever := make(chan bool)

	// Start a goroutine to process messages
	go func() {
		for d := range msgs {
			log.Printf(" [x] Received: %s", d.Body)
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever // Keep the program running
}
