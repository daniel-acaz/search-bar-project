package main

import (
	"github.com/streadway/amqp"
	"search-bar-project/adapters/queue"
	"search-bar-project/application/services"
)

func main() {

	messageChannel := make(chan amqp.Delivery)

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	manager := services.Manager{
		MessageChannel: messageChannel,
		RabbitMQ:       rabbitMQ,
		Result: 		make(chan services.WorkerResult),
	}

	manager.Start(ch)
}
