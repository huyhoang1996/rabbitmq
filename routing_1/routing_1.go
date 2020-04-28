package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	Receive_2()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Receive_2() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.212:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create if queue not exist
	_, err = ch.QueueDeclare(
		"routing_1", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		"routing_1", // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")
	// Ignore if exist
	err = ch.ExchangeDeclare(
		"ex_direct", // name
		"direct",    // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	fmt.Println("==== ExchangeDeclare err", err)
	err = ch.QueueBind(
		"routing_1", // queue name
		"bind_1",    // routing key
		"ex_direct", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	err = ch.QueueBind(
		"routing_1", // queue name
		"bind_all",  // routing key
		"ex_direct", // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			log.Printf("time.Sleep:: ", dot_count)

			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
