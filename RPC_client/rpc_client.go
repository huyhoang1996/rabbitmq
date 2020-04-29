package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	Sender()
}

// CorrelationId

func Sender() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.212:5672/")
	fmt.Println("====  err", err)

	defer conn.Close()

	ch, err := conn.Channel()
	fmt.Println("====  err", err)

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_client", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	body := "Hello World! ha huy haong 2"
	err = ch.Publish(
		"",           // exchange
		"rpc_server", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.Name,
			Body:        []byte(body),
		})

	msgs, err := ch.Consume(
		"rpc_client", // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
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
