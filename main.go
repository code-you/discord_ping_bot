package main

import (
	"fmt"

	"github.com/code-you/discord-ping/bot"
	"github.com/code-you/discord-ping/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config %v:", err)
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}