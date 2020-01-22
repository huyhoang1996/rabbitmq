package main

import (
	"log"

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

func Sender() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.200:5672/")
	failOnErrorSender(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErrorSender(err, "Failed to open a channel")
	defer ch.Close()

	body := "Hello World! ha huy haong 2"
	err = ch.Publish(
		"",       // exchange
		"hello4", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnErrorSender(err, "Failed to publish a message")

}
