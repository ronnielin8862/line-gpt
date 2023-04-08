package global

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/config"
	"log"
	"net/http"
)

var Bot *linebot.Client

func LineInit() {
	client := &http.Client{}
	newBot, err := linebot.New(config.GetConfig().LineServer.ChannelSecret, config.GetConfig().LineServer.ChannelToken, linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatal("init line Bot err :", err)
	}
	Bot = newBot
}
