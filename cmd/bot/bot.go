package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/arl/statsviz"
	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/commands"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/events"
	"github.com/nupamore/pamo_bot/services"
)

func main() {
	if configs.Env["DEBUG_PORT"] != "" {
		go func() {
			statsviz.RegisterDefault()
			log.Fatal(http.ListenAndServe(configs.Env["DEBUG_PORT"], nil))
		}()
	}
	services.AWSsetup()
	services.DBsetup()
	services.Guild.BotStart()

	token := configs.Env["BOT_TOKEN"]
	prefix := configs.Env["BOT_PREFIX"]

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

		// set activity
		ctx.Gateway.Identifier.IdentifyData = gateway.IdentifyData{
			Token: token,
			Presence: &gateway.UpdateStatusData{
				Game: &discord.Activity{
					Name: configs.Env["BOT_STATUS"],
					Type: discord.GameActivity,
				},
				Status: discord.OnlineStatus,
			},
		}

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
