package msgHandler

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/global"
	"line-gpt/server/gpt"
	"log"
	"strings"
)

// test
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
	customizedText := msgCustomized(&text)
	if customizedText == "" {
		return
	}
	log.Printf("customizedText: %s\n", customizedText)
	answer := gpt.Talk(customizedText)
	testPushTextMsg(answer, target)
}

func msgCustomized(s *string) (ns string) {
	switch {
	case strings.HasPrefix(*s, "tc "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tc ", "請將以下內容翻譯成中文: \"", 1), "\"")
	case strings.HasPrefix(*s, "tt "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tt ", "請將以下內容翻譯成泰文: \"", 1), "\"")
	case strings.HasPrefix(*s, "tj "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tj ", "請將以下內容翻譯成日文: \"", 1), "\"")
	case strings.HasPrefix(*s, "ai "):
		ns = strings.Replace(*s, "ai ", "", 1)
	}

	return ns
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
