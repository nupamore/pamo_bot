package main

import (
	"log"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/arikawa/v2/state"
	"github.com/nupamore/pamo_bot/configs"
)

func main() {
	token := "Bot NTAyNDUwNDk0MzgwMTc5NDYx.W8h1fA.jlIhLls0UT19gZaTDxoIFAeVKyQ"
	s, err := state.New(token)
	if err != nil {
		log.Fatal(err)
	}

	s.Gateway.AddIntents(gateway.IntentGuildMessageReactions)
	s.Gateway.AddIntents(gateway.IntentGuildMessages)
	s.AddHandler(func(m *gateway.MessageCreateEvent) {
		log.Println("message")
	})
	s.AddHandler(func(r *gateway.MessageReactionAddEvent) {
		log.Println("react")
	})

	// set activity
	s.Gateway.Identifier.IdentifyData = gateway.IdentifyData{
		Token: token,
		Presence: &gateway.UpdateStatusData{
			Activities: []discord.Activity{{
				Name: configs.Env["BOT_STATUS"],
				Type: discord.GameActivity,
			}},
			Status: gateway.OnlineStatus,
		},
	}

	if err := s.Open(); err != nil {
		log.Fatal(err)
	}

	select {}

	//////////////

	// type Bot struct {
	// 	Ctx *bot.Context
	// }
	// commands := &Bot{}

	// services.AWSsetup()
	// services.DBsetup()
	// services.Guild.BotStart()

	// token := configs.Env["BOT_TOKEN"]
	// prefix := configs.Env["BOT_PREFIX"]

	// services.DiscordAPI = api.NewClient("Bot " + token)

	// wait, err := bot.Start(token, &commands.Commands{}, func(ctx *bot.Context) error {
	// 	ctx.HasPrefix = bot.NewPrefix(prefix)
	// 	me, _ := ctx.Me()

	// 	ctx.AddHandler(events.GuildCreated)
	// 	ctx.AddHandler(events.GuildDeleted)
	// 	ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
	// 		if me.ID == m.Author.ID || m.Author.Bot {
	// 			return
	// 		}
	// 		// no prefix event
	// 		if !strings.HasPrefix(m.Content, prefix) {
	// 			events.NoCommandHandler(m)
	// 		}
	// 	})

	// 	ctx.Gateway.AddIntents(gateway.IntentGuildMessageReactions)
	// 	ctx.Gateway.AddIntents(gateway.IntentGuildMessages)
	// 	ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
	// 		log.Println("message")
	// 	})
	// 	ctx.AddHandler(func(r *gateway.MessageReactionAddEvent) {
	// 		log.Println("react")
	// 	})

	// 	return nil
	// })

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println("Bot started")

	// if err := wait(); err != nil {
	// 	log.Fatalln("Gateway fatal error:", err)
	// }
}
