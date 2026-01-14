package main

import (
	"context"
	"log"

	app "github.com/y2aiskni/prac-go-sacloud-simplemq-api"
)

func main() {
	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	// 初期化
	conf := app.ReadConfig("./config.yaml")
	api := app.NewMessageAPI(conf.Sacloud.Simplemq.QueueName, conf.Sacloud.Simplemq.APIKey)

	// デキュー
	resR, err := api.Receive(ctx)
	if err != nil {
		log.Fatalf("Failed to receive message: %s", err)
	}
	app.OutputJSON(resR)
	if err := api.Delete(ctx, string(resR[0].ID)); err != nil {
		log.Fatalf("Failed to delete message: %s", err)
	}
}
