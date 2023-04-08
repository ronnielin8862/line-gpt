package main

import (
	"line-gpt/config"
	"line-gpt/server/gpt"
	"line-gpt/server/line"
	"log"
)

func main() {
	_, err := config.LoadGlobalConfig()
	if err != nil {
		log.Fatal("init config err :", err)
	}

	gpt.Start()
	line.Init()
}
