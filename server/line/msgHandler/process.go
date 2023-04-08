package msgHandler

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/server/line/lineUtil"
	"log"
)

func ChannelProcessor() {
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
