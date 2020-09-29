package main

import (
	"log"
	"os"
	"strings"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/joho/godotenv"
	"github.com/nupamore/pamo_bot/discord"
	"github.com/nupamore/pamo_bot/services"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	if err := services.DBsetup(); err != nil {
		log.Println("DB init fail")
		panic(err)
	}

	token := os.Getenv("BOT_TOKEN")
	prefix := os.Getenv("BOT_PREFIX")

	commands := &discord.Commands{}

	wait, err := bot.Start(token, commands, func(ctx *bot.Context) error {
		ctx.HasPrefix = bot.NewPrefix(prefix)
		me, _ := ctx.Me()

		ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
			if me.ID == m.Author.ID {
				return
			}
			// no prefix event
			if !strings.HasPrefix(m.Content, prefix) {
				discord.NoCommandHandler(m)
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