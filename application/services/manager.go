package services

import (
	"github.com/streadway/amqp"
	"log"
	"search-bar-project/adapters/queue"
	"search-bar-project/domain"
	"time"
)

type Manager struct {
	MessageChannel 		chan amqp.Delivery
	RabbitMQ 			*queue.RabbitMQ
	Result				chan WorkerResult
}

func (m *Manager) Start(ch *amqp.Channel){

	concurrency := 4

	search := domain.Search{
		ID:         "",
		Avatar:     "",
		Name:       "",
		Categories: nil,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	for qtdProcess := 0; qtdProcess < concurrency; qtdProcess++ {
		go Worker(m.MessageChannel, search)
	}

	for workerResult := range m.Result {
		if workerResult.Error != nil {
			log.Fatalf("[ERROR]: %v", workerResult.Error.Error())
		} else {
			workerResult.Message.Ack(false)
		}

	}
}