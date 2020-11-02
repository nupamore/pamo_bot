package commands

import (
	"fmt"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/services"
)

// Crawl : crawl images command
func (cmd *Commands) Crawl(m *gateway.MessageCreateEvent, arg bot.RawArguments) (string, error) {
	switch arg {
	// current status
	case "":
		_, isScrapingChannel := services.Guild.ScrapingChannelIDs[m.ChannelID]
		if isScrapingChannel {
			return "Watching now this channel", nil
		}
		return "Not watching this channel", nil

	// change status
	case "on":
		services.Guild.AddScrapingChannel(m.GuildID, m.ChannelID)
		return "Watching now this channel", nil
	case "off":
		services.Guild.RemoveScrapingChannel(m.GuildID)
		return "Not watching this channel", nil

	// crawl past
	case "past":
		newImg, max, msgID := 0, 20, m.ID

		temp := "Crawling messages (%d / %d)"
		sentMsg, _ := services.DiscordAPI.SendText(
			m.ChannelID,
			fmt.Sprintf(temp, 0, 100),
		)

		for i := 1; i <= max; i++ {
			n, id, _ := services.Image.Crawl(m.ChannelID, m.GuildID, msgID)
			newImg = newImg + n
			msgID = id

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
		return fmt.Sprintf("New images: %d", newImg), nil
	}

	return "", nil
}
