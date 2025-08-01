package main

import (
	"fmt"

	"github.com/pcubedp/discord_pingbot/bot"
	"github.com/pcubedp/discord_pingbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
