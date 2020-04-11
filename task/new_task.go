package main

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	Sender()
}

func failOnErrorSender(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func Sender() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.1.2:5672/")
	failOnErrorSender(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErrorSender(err, "Failed to open a channel")
	defer ch.Close()

	body := bodyFrom([]string{"ha ..", "huy ..."})
	err = ch.Publish(
		"",         // exchange
		"new_task", // routing key
		false,      // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // chua hieu
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnErrorSender(err, "Failed to publish a message")

}
