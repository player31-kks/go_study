package main

import (
	"listener/event"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	connection, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer connection.Close()

	// start listening for messages
	log.Println("Listening for and consuming messages...")

	// create consumer
	consumer, err := event.NewConsumer(connection)
	if err != nil {
		log.Println(err)
	}

	// watch the queue and conusume events
	err = consumer.Listen([]string{"log.INFO", "log.ERROR", "log.WARNING"})
	if err != nil {
		log.Println(err)
	}

}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	for {
		var err error
		connection, err = amqp.Dial("amqp://rabbitmq:password@rabbitmq:5672/")
		if err == nil {
			return connection, nil
		}

		counts++
		time.Sleep(backoff)
		backoff = backoff * 2
		if backoff > 60*time.Second {
			backoff = 60 * time.Second
		}
	}

}
