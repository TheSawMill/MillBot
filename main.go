package main

import memes "discord-bot/memes"
import bot "discord-bot/bot"
import config "discord-bot/config"

import (
	"log"
)

func main() {
	configuration := config.GetConfiguration()

	if configuration == nil {
		log.Fatalln("Configuration is nil exiting program.")
	}

	memes.TestGetRequest()
	bot.BotToken = configuration.BotToken
	bot.Run() // call the run function of bot/bot.go
}
