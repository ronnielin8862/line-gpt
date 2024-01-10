package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"line-gpt/global"
	"log"
	"strings"
)

func ImageCreate(ask string) (urls []string) {
	ask = strings.Replace(ask, "ci ", "", 1)
	log.Printf("ask: %s\n", ask)
	resp, err := global.GptClient.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Model:          openai.CreateImageModelDallE3,
			Size:           openai.CreateImageSize1024x1024,
			Prompt:         ask,
			N:              1,
			Quality:        openai.CreateImageQualityStandard,
			ResponseFormat: openai.CreateImageResponseFormatURL,
		},
	)
	if err != nil {
		errString := fmt.Sprintf("ImageCreate error: %v", err)
		log.Printf("ImageCreate error: %v\n\n", errString)
		urls = append(urls, "有什麼錯誤發生了！ : %s", errString)
		return urls
	}
	for _, c := range resp.Data {
		urls = append(urls, c.URL)
	}
	log.Println(urls)
	return urls
}
