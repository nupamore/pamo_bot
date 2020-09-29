package discord

import (
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/services"
)

// Commands : has prefix
type Commands struct {
	Ctx *bot.Context
}

// Service : Service
var Service *services.Service

// NoCommandHandler : if has not prefix
func NoCommandHandler(m *gateway.MessageCreateEvent) {
	// log.Println(m.Author.Username, "sent", m.Content)
}
