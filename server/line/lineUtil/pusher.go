package lineUtil

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/global"
	"log"
)

var TextChannel = make(chan *linebot.Event, 10)
var ImageChannel = make(chan *linebot.Event, 3)

func CheckTextChannelSize(event *linebot.Event) {
	log.Println("textChannel size: ", len(TextChannel))
	if len(TextChannel) >= 9 {
		PushTextMsg("目前等待回覆問題量較大，等一分鐘後再問辣...", event)
	} else if len(TextChannel) >= 2 {
		PushTextMsg("稍等下", event)
	}
}

func CheckImgChannelSize(event *linebot.Event) {
	log.Println("imgChannel size: ", len(ImageChannel))
	if len(ImageChannel) >= 5 {
		PushTextMsg("目前圖片處理量較大，等一分鐘後再問辣...", event)
	} else if len(ImageChannel) >= 1 {
		PushTextMsg("稍等下", event)
	}
}

func getTarget(event *linebot.Event) string {
	userID := event.Source.UserID
	groupID := event.Source.GroupID
	RoomID := event.Source.RoomID

	var target string
	if RoomID != "" {
		target = RoomID
	} else if groupID != "" {
		target = groupID
	} else {
		target = userID
	}

	log.Println("RoomID: ", RoomID, ", groupID: ", groupID, ", userID: ", userID)
	return target
}

func PushTextMsg(content string, event *linebot.Event) {

	target := getTarget(event)

	var messages []linebot.SendingMessage
	messages = append(messages, linebot.NewTextMessage(content))
	response, err := global.Bot.PushMessage(target, messages...).Do()
	if err != nil {
		log.Println(err)
	}

	log.Printf("Push Msg: %s, Response Id: %s", content, response.RequestID)
}

func PushImageMsg(url string, event *linebot.Event) {

	target := getTarget(event)
	log.Println("PushImageMsg: ", url, ", target: ", target)
	var messages []linebot.SendingMessage
	messages = append(messages, linebot.NewImageMessage(url, url))
	response, err := global.Bot.PushMessage(target, messages...).Do()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Push Msg: %s, Response Id: %s", url, response.RequestID)
	}

}
