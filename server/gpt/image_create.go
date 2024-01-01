package gpt

import (
	"context"
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
			Model:          openai.GPT4TurboPreview,
			Size:           openai.CreateImageSize512x512,
			Prompt:         ask,
			N:              2,
			ResponseFormat: openai.CreateImageResponseFormatURL,
		},
	)
	if err != nil {
		log.Printf("ImageCreate error: %v\n\n", err)
		urls = append(urls, "我有點錯亂，請再試一次...")
		return
	}
	for _, c := range resp.Data {
		urls = append(urls, c.URL)
	}
	log.Println(urls)
	return urls
}
