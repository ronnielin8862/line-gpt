package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"line-gpt/global"
	"log"
)

func Talk(ask string) string {
	log.Printf("ask: %s\n", ask)
	resp, err := global.GptClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4TurboPreview,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: ask,
				},
			},
		},
	)

	if err != nil {
		errString := fmt.Sprintf("ImageCreate error: %v", err)
		log.Printf("ImageCreate error: %v\n\n", errString)
		return fmt.Sprintf("ImageCreate error: %v\n\n", errString)
	}

	log.Println(resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content
}
