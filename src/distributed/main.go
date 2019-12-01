package distributed

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func mssgClient() {

	conn, ch, qu := getQueue()

	defer conn.Close()
	defer ch.Close()

	mssgs, err := ch.Consume(
		qu.Name,
		"", // Name of consuming client - if tracking
		true, // auto acknowledge receipt of message, and so remove from queue
		false, // True if client should be exclusive consumer of this queue
		false, // disables local host
		false,
		nil,
		)

	failOnError(err, "Failed to receive messages.")

	for mssg := range mssgs {
		fmt.Println("This is the message: ", mssg.Body)
	}

}

// Send messages to rabbit
func mssgServer() {
	// Get connection to queue
	conn, ch, qu := getQueue()
	// At the end, close channel and queue
	defer conn.Close()
	defer ch.Close()

	// Define message struct to send to rabbit exchange
	mssg := amqp.Publishing{
		ContentType:     "text/plain",
		Timestamp:       time.Time{},
		Body:            []byte("Hello world!"),
	}

	// will keep sending messages until programme forcibly exited
	for {
		err := ch.Publish("", qu.Name, false, false, mssg)
		failOnError(err, "Failed to publish message to rabbit.")

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





func main() {
	// Message server and client run as go routines to allow continuous execution
	go mssgServer()
	go mssgServer()

	// keep main go routine running so that we can obserce the others executing
	var st string
	fmt.Scanln(&st)

}