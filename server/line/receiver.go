package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/global"
	"line-gpt/server/line/msgHandler"
	"log"
	"net/http"
)

func receiver(w http.ResponseWriter, req *http.Request) {
	log.Println("x-line-signature: ", req.Header.Get("x-line-signature"))

	events, err := global.Bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Println("invalid signature: ", err)
			w.WriteHeader(400)
		} else {
			log.Println("other error: ", err)
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			j, _ := event.MarshalJSON()
			log.Println("Receive EventType = Message, content : ", string(j))

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msgHandler.TextMsgHandler(message.Text, event)
			default:
				return
			}

		} else {
			log.Println("Received Other EventType: ", event.Message)
		}

	}
}
