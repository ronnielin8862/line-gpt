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
		errString := fmt.Sprintf("talk error: %v", err)
		log.Printf("%v\n\n", errString)
		return errString
	}

	var response = resp.Choices[0].Message.Content
	log.Println("ask answer " + response)
	log.Println("ask answer " + response)
	return response
}
