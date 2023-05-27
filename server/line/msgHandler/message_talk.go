package msgHandler

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/server/gpt"
	"line-gpt/server/line/lineUtil"
	"log"
	"strings"
	"time"
)

func TextMsgHandler(text string, event *linebot.Event) {
	switch processType(text) {
	case "chat":
		lineUtil.TextChannel <- event
		lineUtil.CheckTextChannelSize(event)
	case "imgCreate":
		log.Println("imgCreate 1 ")
		lineUtil.ImageChannel <- event
		lineUtil.CheckImgChannelSize(event)
	default:
		return
	}
}

func processType(s string) (ns string) {
	switch {
	case strings.HasPrefix(s, "tc "):
		ns = "chat"
	case strings.HasPrefix(s, "tt "):
		ns = "chat"
	case strings.HasPrefix(s, "tj "):
		ns = "chat"
	case strings.HasPrefix(s, "te "):
		ns = "chat"
	case strings.HasPrefix(s, "tk "):
		ns = "chat"
	case strings.HasPrefix(s, "ci "):
		ns = "imgCreate"
	case strings.HasPrefix(s, "ai "):
		ns = "chat"
	}
	return ns
}

func textProcess(content string, event *linebot.Event) {
	customizedText := msgCustomized(&content)
	answer := gpt.Talk(customizedText)
	lineUtil.PushTextMsg(answer, event)
}

func msgCustomized(s *string) (ns string) {
	ss := *s
	if len(ss) >= 4 {
		ss = strings.ToLower(ss[:3]) + ss[3:]
	}
	switch {
	case strings.HasPrefix(ss, "tc "):
		ns = fmt.Sprintf("%s%s", strings.Replace(ss, "tc ", "請將以下內容翻譯成中文: \"", 1), "\"")
	case strings.HasPrefix(ss, "tt "):
		ns = fmt.Sprintf("%s%s", strings.Replace(ss, "tt ", "請將以下內容翻譯成泰文: \"", 1), "\"")
	case strings.HasPrefix(ss, "tj "):
		ns = fmt.Sprintf("%s%s", strings.Replace(ss, "tj ", "請將以下內容翻譯成日文: \"", 1), "\"")
	case strings.HasPrefix(ss, "te "):
		ns = fmt.Sprintf("%s%s", strings.Replace(ss, "te ", "請將以下內容翻譯成英文: \"", 1), "\"")
	case strings.HasPrefix(ss, "tk "):
		ns = fmt.Sprintf("%s%s", strings.Replace(ss, "tk ", "請將以下內容翻譯成高棉文: \"", 1), "\"")
	case strings.HasPrefix(ss, "ai "):
		ns = strings.Replace(ss, "ai ", "", 1)
	}

	return ns
}

func imageProcess(content string, event *linebot.Event) {
	log.Printf("imageProcess 1 %s", content)
	urls := gpt.ImageCreate(content)
	for _, url := range urls {
		lineUtil.PushImageMsg(url, event)
		time.Sleep(1 * time.Second)
	}
}
