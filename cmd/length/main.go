package main

import (
	"context"
	"fmt"
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
	api := app.NewQueueAPI(conf.Sacloud.AccessToken, conf.Sacloud.AccessTokenSecret)

	// 長さの取得
	c, err := api.CountMessages(ctx, conf.Sacloud.Simplemq.ResourceID)
	if err != nil {
		log.Fatalf("Failed to get message count: %s", err)
	}
	fmt.Println(c)

}
