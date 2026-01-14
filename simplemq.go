package app

import (
	"log"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/simplemq-api-go"
)

func NewQueueAPI(accessToken string, accessTokenSecret string) simplemq.QueueAPI {
	client, err := simplemq.NewQueueClient(client.WithApiKeys(accessToken, accessTokenSecret))
	if err != nil {
		log.Fatalf("Failed to create SimpleMQ client: %s", err)
	}
	return simplemq.NewQueueOp(client)
}

func NewMessageAPI(queueName string, apiKey string) simplemq.MessageAPI {
	client, err := simplemq.NewMessageClient(apiKey)
	if err != nil {
		log.Fatalf("Failed to create SimpleMQ client: %s", err)
	}
	return simplemq.NewMessageOp(client, queueName)
}
