package discord

import (
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
)

// Commands : has prefix
type Commands struct {
	Ctx *bot.Context
}

// NoCommandHandler : if has not prefix
func NoCommandHandler(m *gateway.MessageCreateEvent) {
	// log.Println(m.Author.Username, "sent", m.Content)
}
