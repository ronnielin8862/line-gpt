package lineUtil

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/global"
	"log"
)

var textChannel = make(chan *linebot.Event, 10)

func CheckChannelSize(event *linebot.Event) {
	if len(textChannel) >= 9 {
		PushTextMsg("目前等待回覆問題量較大，等一分鐘後再問辣...", event)
	}
	if len(textChannel) >= 2 {
		PushTextMsg("稍等下", event)
	}
}

func PushTextMsg(content string, event *linebot.Event) {

	userID := event.Source.UserID
	groupID := event.Source.GroupID
	RoomID := event.Source.RoomID

	var target string
	if RoomID != "" {
		target = RoomID
		goto ASK
	} else if groupID != "" {
		target = groupID
		goto ASK
	} else {
		target = userID
	}

ASK:

	var messages []linebot.SendingMessage
	messages = append(messages, linebot.NewTextMessage(content))
	response, err := global.Bot.PushMessage(target, messages...).Do()
	if err != nil {
		log.Println(err)
	}

	log.Println("RoomID: ", RoomID, ", groupID: ", groupID, ", userID: ", userID)
	log.Printf("Push Msg: %s, Response Id: %s", content, response.RequestID)
}
