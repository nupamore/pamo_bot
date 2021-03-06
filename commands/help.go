package commands

import (
	"strings"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/nupamore/pamo_bot/configs"
)

// Help : descriptions
func (cmd *Commands) Help(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	prefix := configs.Env["BOT_PREFIX"]
	desc := `
**$t [target] [text]**
Translate any [text] to [target] language
**$dice [max]**
Get a random number. [max] is the maximum
**$crawl [on/off]**
Activate real-time image scraping in this channel
**$crawl past**
Scraping past images
**$image [username]**
Get a random image uploaded by [username]
    `
	description := &discord.Embed{
		Title:       "Pamo_bot commands list",
		Description: strings.Replace(desc, "$", prefix, -1),
		Fields: []discord.EmbedField{
			{Name: "**Photo Archive**", Value: "https://bot.nupa.moe"},
		},
	}

	return description, nil
}
