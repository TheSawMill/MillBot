package main

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

	bot.BotToken = configuration.BotToken
	bot.GifConfigurationKey = configuration.GifyApiKey
	bot.Run() // call the run function of bot/bot.go
}
