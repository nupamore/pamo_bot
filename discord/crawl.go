package discord

import (
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/services"
)

// Crawl : crawl images command
func (cmd *Commands) Crawl(m *gateway.MessageCreateEvent, arg bot.RawArguments) (string, error) {
	// current status
	if arg == "" {
		_, isScrapingChannel := services.ScrapingChannelIDs[m.ChannelID]
		if isScrapingChannel {
			return "Watching now this channel", nil
		}
		return "Not watching this channel", nil
	}

	if arg == "on" {
		services.AddScrapingChannel(m.GuildID, m.ChannelID)
		return "Watching now this channel", nil
	}

	if arg == "off" {
		services.RemoveScrapingChannel(m.GuildID, m.ChannelID)
		return "Not watching this channel", nil
	}

	return "", nil
}
