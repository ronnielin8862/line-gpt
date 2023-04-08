package line

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/global"
	"line-gpt/server/line/msgHandler"
	"net/http"
)

func testReceive(w http.ResponseWriter, req *http.Request) {
	fmt.Println("x-line-signature: ", req.Header.Get("x-line-signature"))

	events, err := global.Bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			fmt.Println("invalid signature: ", err)
			w.WriteHeader(400)
		} else {
			fmt.Println("other error: ", err)
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		fmt.Println("received event Type : ", event.Type)
		if event.Type == linebot.EventTypeMessage {
			fmt.Println("received EventTypeMessage")
			j, _ := event.MarshalJSON()
			fmt.Println(string(j))
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(message.Text)
				msgHandler.TextMsg(message.Text, event)
			}

		} else {
			fmt.Println("received other: ", event.Message)
		}

	}
}
