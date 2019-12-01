package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)


func mssgServer() {
	// Get connection to queue
	conn, ch, qu := getQueue()
	// At the end, close channel and queue
	defer conn.Close()
	defer ch.Close()

	// Define message struct to send to rabbit exchange
	mssg := amqp.Publishing{
		Headers:         nil,
		ContentType:     "",
		ContentEncoding: "",
		DeliveryMode:    0,
		Priority:        0,
		CorrelationId:   "",
		ReplyTo:         "",
		Expiration:      "",
		MessageId:       "",
		Timestamp:       time.Time{},
		Type:            "",
		UserId:          "",
		AppId:           "",
		Body:            nil,
	}
}


func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {

	conn, err := amqp.Dial("amqp://guest@localhost:5672")

	// If connection has failed, clean up
	failOnError(err, "Connection to rabbitMQ failed.")

	ch, err := conn.Channel()

	failOnError(err, "Failed to open new channel.")

	qu, err := ch.QueueDeclare("new_queue",
		false, // don't persist messages to hard disk
		false, // keep messages in queue until received
		false, // non-exclusive queue
		false,
		nil, // no args required
		)

	failOnError(err, "Failed to declare queue.")

	return conn, ch, &qu

}

// If there is an error connecting to rabbit mq, exit app
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)

		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}





func main {


}