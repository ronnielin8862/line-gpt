package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
)

var bot *linebot.Client

func main() {
	client := &http.Client{}
	newBot, err := linebot.New("15c23eec6d4818bbbc7ff242f1392bbf", "cnLFx4O7KJa4pzIKKKfHWuDZLBExpEHcRiCNSInzirMPaST+2+pntYZ3ANayuCcGyzHTsgQW1ZHezGCEYxiKnwAJhFvwJilrjWGHOaGdhEH0pboysT1o7yhAhTYc0EpkUUFV0DaCg7qEZykq8QoA1wdB04t89/1O/w1cDnyilFU=", linebot.WithHTTPClient(client))
	//newBot, err := linebot.New("15c23eec6d4818bbbc7ff242f1392bbf", "cnLFx4O7KJa4pzIKKKfHWuDZLBExpEHcRiCNSInzirMPaST+2+pntYZ3ANayuCcGyzHTsgQW1ZHezGCEYxiKnwAJhFvwJilrjWGHOaGdhEH0pboysT1o7yhAhTYc0EpkUUFV0DaCg7qEZykq8QoA1wdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		log.Fatal(err)
	}
	bot = newBot
	testPushTextMsg()
	//testPushActionMsg()
	router := mux.NewRouter()
	router.HandleFunc("/test", testReceive).Methods("POST")
	log.Fatal(http.ListenAndServe(":9487", router))
}

func testPushTextMsg() {
	var messages []linebot.SendingMessage
	messages = append(messages, linebot.NewTextMessage("Hello World"))
	response, err := bot.PushMessage("U6231183579419303e37e265b8da4e2a1", messages...).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response: ", response.RequestID)
}

func testPushActionMsg() {
	leftBtn := linebot.NewMessageAction("left", "left clicked")
	rightBtn := linebot.NewMessageAction("right", "right clicked")

	template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

	message := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)

	var messages []linebot.SendingMessage
	messages = append(messages, message)
	response, err := bot.PushMessage("U6231183579419303e37e265b8da4e2a1", messages...).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response: ", response.RequestID)
}

func testReceive(w http.ResponseWriter, req *http.Request) {
	fmt.Println("x-line-signature: ", req.Header.Get("x-line-signature"))

	//body, err := io.ReadAll(req.Body)
	//if err != nil {
	//	fmt.Println("read body error: ", err)
	//	return
	//}
	//fmt.Println("req.body: ", string(body))

	//return

	events, err := bot.ParseRequest(req)
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

			}

			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID

			fmt.Println("userID: ", userID, " groupID: ", groupID, " RoomID: ", RoomID)
		}else {
			fmt.Println("received other: ", event.Message)
		}

	}
}