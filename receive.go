package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Receive() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.1.2:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create if queue not exist
	_, err = ch.QueueDeclare(
		"hello3", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		"hello3", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	fmt.Println("====== ")
	go func() {
		fmt.Println("Inside routine ")

		for d := range msgs {
			fmt.Println("Inside routine FOR ")
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("====== ")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
