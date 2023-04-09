package gpt

import (
	"line-gpt/config"
	"testing"
)

func TestImageCreate(t *testing.T) {
	config.LoadGlobalConfig()
	Start()

	ask := "我想要貓站在高樓陽台邊的圖片"
	urls := ImageCreate(ask)
	t.Log(urls)
}
