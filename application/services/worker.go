package services

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"log"
	"search-bar-project/application/repositories"
	"search-bar-project/domain"
)

type WorkerResult struct {
	Message 		*amqp.Delivery
	Error			error
}

func Worker(messageChannel chan amqp.Delivery, search domain.Search) {

	for message := range messageChannel {

		err := json.Unmarshal(message.Body, &search)
		search.ID = uuid.NewV4().String()


		if err != nil {
			log.Fatalf("[ERROR]: %v", err.Error())
			continue
		}

		repositories.Save(search)
		log.Printf("The name of the search result is: %s", search.Name)
	}

}