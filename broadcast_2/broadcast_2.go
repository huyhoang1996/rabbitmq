package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	Broadcast()
}
func failOnErrorSender(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func Broadcast() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.212:5672/")
	fmt.Println("====  err", err)

	failOnErrorSender(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	fmt.Println("====  err", err)

	failOnErrorSender(err, "Failed to open a channel")
	defer ch.Close()
	fmt.Println("====  err", err)

	err = ch.ExchangeDeclare(
		"ex_broadcast", // name
		"fanout",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	fmt.Println("==== ExchangeDeclare err", err)

	body := "HOang buon 2222"
	err = ch.Publish(
		"ex_broadcast", // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	fmt.Println("==== Publish err", err)

}
