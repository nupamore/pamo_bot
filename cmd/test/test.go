package main

import (
	"github.com/diamondburned/arikawa/api"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/services"
)

func main() {
	services.DiscordAPI = api.NewClient("Bot " + configs.Env["BOT_TOKEN"])
	services.SendDM(1234, "dm test")
}
