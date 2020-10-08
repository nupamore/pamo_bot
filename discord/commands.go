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

// NoCommandHandler : if has not prefix
func NoCommandHandler(m *gateway.MessageCreateEvent) {
	// scrap image
	hasImage := len(m.Attachments) > 0
	_, isScrapingChannel := services.ScrapingChannelIDs[m.ChannelID]

	if hasImage && isScrapingChannel {
		services.ScrapImage(m.Attachments[0])
	}
}
