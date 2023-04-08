package gpt

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"line-gpt/global"
	"log"
)

func Talk(ask string) string {
	log.Printf("ask: %s\n", ask)
	resp, err := global.GptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: ask,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n\n", err)
		return "我有點錯亂，請再問一次..."
	}

	log.Println(resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content
}
