package main

import (
	"log"
	"os"
	"strings"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/joho/godotenv"
	"github.com/nupamore/pamo_bot/commands"
	"github.com/nupamore/pamo_bot/events"
	"github.com/nupamore/pamo_bot/services"
)

func init() {
	godotenv.Load("configs/.env")
}

func main() {
	services.AWSsetup()
	services.DBsetup()
	services.InitGuildsInfo()

	token := os.Getenv("BOT_TOKEN")
	prefix := os.Getenv("BOT_PREFIX")

	services.DiscordAPI = api.NewClient("Bot " + token)

	wait, err := bot.Start(token, &commands.Commands{}, func(ctx *bot.Context) error {
		ctx.HasPrefix = bot.NewPrefix(prefix)
		me, _ := ctx.Me()

		ctx.AddHandler(events.GuildCreated)
		ctx.AddHandler(events.GuildDeleted)
		ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
			if me.ID == m.Author.ID || m.Author.Bot {
				return
			}
			// no prefix event
			if !strings.HasPrefix(m.Content, prefix) {
				events.NoCommandHandler(m)
			}
		})

		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Bot started")

	if err := wait(); err != nil {
		log.Fatalln("Gateway fatal error:", err)
	}
}
