package discord

import (
	"fmt"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
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

	// change status
	if arg == "on" {
		services.AddScrapingChannel(m.GuildID, m.ChannelID)
		return "Watching now this channel", nil
	}
	if arg == "off" {
		services.RemoveScrapingChannel(m.GuildID, m.ChannelID)
		return "Not watching this channel", nil
	}

	// crawl past
	if arg == "past" {
		max, msgID := 20, m.ID

		temp := "Crawling messages (%d / %d)"
		sentMsg, _ := services.DiscordAPI.SendText(
			m.ChannelID,
			fmt.Sprintf(temp, 0, 100),
		)

		for i := 1; i <= max; i++ {
			msgID, _ = services.CrawlImages(m.ChannelID, msgID)

			if msgID == discord.NullMessageID {
				i = max
			}

			services.DiscordAPI.EditText(
				m.ChannelID,
				sentMsg.ID,
				fmt.Sprintf(temp, int(
					(float32(i)/float32(max))*100),
					100,
				),
			)
		}
		return "Done!", nil
	}

	return "", nil
}
