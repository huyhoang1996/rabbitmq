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
		"topic_1", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		"topic_1", // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")
	// Ignore if exist
	err = ch.ExchangeDeclare(
		"ex_topic", // name
		"topic",    // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	fmt.Println("==== ExchangeDeclare err", err)
	err = ch.QueueBind(
		"topic_1",   // queue name
		"*.firsh.#", // routing key
		"ex_topic",  // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	err = ch.QueueBind(
		"topic_1",  // queue name
		"eat.#",    // routing key
		"ex_topic", // exchange
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
