package main

import (
	"line-gpt/config"
	"line-gpt/server/gpt"
	"line-gpt/server/line"
)

func main() {
	config.LoadGlobalConfig()
	gpt.Start()
	line.Init()
}
