package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// func main() {
// 	Sender()
// }

// func failOnErrorSender(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

func Sender() {

	conn, err := amqp.Dial("amqp://admin:admin@192.168.3.212:5672/")
	fmt.Println("====  err", err)

	defer conn.Close()

	ch, err := conn.Channel()
	fmt.Println("====  err", err)

	defer ch.Close()

	body := "Hello World! ha huy haong 2"
	err = ch.Publish(
		"",          // exchange
		"broadcast", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

}
