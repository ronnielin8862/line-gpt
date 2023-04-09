package gpt

import (
	"line-gpt/config"
	"testing"
)

func TestImageCreate(t *testing.T) {
	config.LoadGlobalConfig()
	Start()

	ask := "臘腸狗"
	urls := ImageCreate(ask)
	t.Log(urls)
}
