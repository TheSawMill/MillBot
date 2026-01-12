package bot

import memes "discord-bot/memes"

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string
var GifConfigurationKey string
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "hello",
		Description: "Replies with Hello World",
	},
	{
		Name:        "random-meme",
		Description: "Gets a random gif meme",
	},
}

func checkNilError(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}

func Run() {

	// create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilError(err)

	discord.Identify.Intents =
		discordgo.IntentsGuilds |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsMessageContent

	registerCommands(discord)

	// add a event handler
	discord.AddHandler(newMessage)

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		switch i.ApplicationCommandData().Name {
		case "hello":
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello world",
				},
			})
			if err != nil {
				log.Println("Interaction response error:", err)
			}
		case "random-meme":
			url, err := memes.GetRandomMeme(GifConfigurationKey)
			if err != nil {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Failed to fetch meme.",
					},
				})
				return
			}
			log.Println("Url:", url)
			// Respond with the GIF
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: url,
				},
			})

		}
	})

	// open a session
	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open Discord session: %v", err)
	}

	defer discord.Close() // close session, after function termination

	// keep bot running unit there is NO os interruption (ctrl + C)
	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	/* prevent bot responding to its own message
	this is achieved by looking into the message author id
	if the message.author.id is same as bot.author.id then just return
	*/
	if message.Author.ID == discord.State.User.ID {
		return
	}

	log.Println("Message Received from Discord")
	log.Println("From: " + message.Author.ID)
	log.Println("Message: " + message.Content)

	switch {
	case strings.Contains(message.Content, "!help"):
		discord.ChannelMessageSend(message.ChannelID, "Hello World!")
	case strings.Contains(message.Content, "!bye"):
		discord.ChannelMessageSend(message.ChannelID, "Good bye!")
	}
}

func registerCommands(discord *discordgo.Session) {
	u, err := discord.User("@me")
	if err != nil {
		log.Fatalf("Cannot get bot user: %v", err)
	}

	appId := u.ID

	for _, cmd := range commands {
		_, err := discord.ApplicationCommandCreate(appId, "", cmd)
		if err != nil {
			log.Fatalf("Cannot create '%s' command: %v", cmd.Name, err)
		}
	}
}
