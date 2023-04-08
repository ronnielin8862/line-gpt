package msgHandler

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/global"
	"line-gpt/server/gpt"
	"log"
)

func TextMsg(text string, event *linebot.Event) {

	userID := event.Source.UserID
	groupID := event.Source.GroupID
	RoomID := event.Source.RoomID
	fmt.Println("userID: ", userID, " groupID: ", groupID, " RoomID: ", RoomID)
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
	answer := gpt.Talk(text)
	testPushTextMsg(answer, target)
}

func testPushTextMsg(answer, target string) {
	var messages []linebot.SendingMessage
	messages = append(messages, linebot.NewTextMessage(answer))
	response, err := global.Bot.PushMessage(target, messages...).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response: ", response.RequestID)
}
