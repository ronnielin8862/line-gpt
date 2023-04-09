package msgHandler

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/server/line/lineUtil"
	"log"
)

func TextChannelProcessor() {
	for {
		event := <-lineUtil.TextChannel
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			textProcess(message.Text, event)
		default:
			log.Println("process another event type :  , do nothing! ", event.Type)
		}
	}
}

func ImageChannelProcessor() {
	for {
		event := <-lineUtil.ImageChannel
		log.Println("ImageChannelProcessor 1 ")
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			imageProcess(message.Text, event)
		default:
			log.Println("process another event type :  , do nothing! ", event.Type)
		}
	}
}
