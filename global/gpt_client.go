package global

import (
	"github.com/sashabaranov/go-openai"
	"line-gpt/config"
)

var GptClient *openai.Client

func GptInit() {
	GptClient = openai.NewClient(config.GetConfig().GptConf.AuthToken)
}
