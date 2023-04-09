package msgHandler

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-gpt/server/gpt"
	"line-gpt/server/line/lineUtil"
	"strings"
)

func TextMsgHandler(text string, event *linebot.Event) {
	switch processType(&text) {
	case "chat":
		lineUtil.TextChannel <- event
		lineUtil.CheckTextChannelSize(event)
	case "imgCreate":
		lineUtil.ImageChannel <- event
		lineUtil.CheckImgChannelSize(event)
	default:
		return
	}
}

func processType(s *string) (ns string) {
	switch {
	case strings.HasPrefix(*s, "tc "):
		ns = "chat"
	case strings.HasPrefix(*s, "tt "):
		ns = "chat"
	case strings.HasPrefix(*s, "tj "):
		ns = "chat"
	case strings.HasPrefix(*s, "te "):
		ns = "chat"
	case strings.HasPrefix(*s, "tk "):
		ns = "chat"
	case strings.HasPrefix(*s, "ci "):
		ns = "imgCreate"
	case strings.HasPrefix(*s, "ai "):
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
	switch {
	case strings.HasPrefix(*s, "tc "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tc ", "請將以下內容翻譯成中文: \"", 1), "\"")
	case strings.HasPrefix(*s, "tt "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tt ", "請將以下內容翻譯成泰文: \"", 1), "\"")
	case strings.HasPrefix(*s, "tj "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tj ", "請將以下內容翻譯成日文: \"", 1), "\"")
	case strings.HasPrefix(*s, "te "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "te ", "請將以下內容翻譯成英文: \"", 1), "\"")
	case strings.HasPrefix(*s, "tk "):
		ns = fmt.Sprintf("%s%s", strings.Replace(*s, "tk ", "請將以下內容翻譯成高棉文: \"", 1), "\"")
	case strings.HasPrefix(*s, "ai "):
		ns = strings.Replace(*s, "ai ", "", 1)
	}

	return ns
}

func imageProcess(content string, event *linebot.Event) {
	urls := gpt.ImageCreate(content)
	for _, url := range urls {
		lineUtil.PushTextMsg(url, event)
	}
}

func imgCustomized(s *string) (ns string) {
	switch {
	case strings.HasPrefix(*s, "ci "):
		ns = strings.Replace(*s, "ci ", "", 1)
	}
	return ns
}
