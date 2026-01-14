package main

import (
	"context"
	"log"
	"os"

	"github.com/sacloud/simplemq-api-go/apis/v1/message"
	app "github.com/y2aiskni/prac-go-sacloud-simplemq-api"
)

func main() {
	ctx := context.Background()
	run(ctx, os.Args[1:]...)
}

func run(ctx context.Context, messages ...string) {
	if len(messages) == 0 {
		log.Println("No messages to enqueue")
		return
	}

	// 初期化
	conf := app.ReadConfig("./config.yaml")
	api := app.NewMessageAPI(conf.Sacloud.Simplemq.QueueName, conf.Sacloud.Simplemq.APIKey)

	// エンキュー
	var result []*message.NewMessage
	for _, msg := range messages {
		res, err := api.Send(ctx, msg)
		if err != nil {
			log.Fatalf("Failed to enqueue \"%s\" message: %s", msg, err)
		}
		result = append(result, res)
	}
	app.OutputJSON(result)
}
