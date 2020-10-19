package discord

import (
	"os"
	"strings"

	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Help : descriptions
func (cmd *Commands) Help(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	prefix := os.Getenv("BOT_PREFIX")
	desc := `
__$t [target] [text]__
Translate any [text] to [target] language
__$Dice [max]__
Get a random number. [max] is the maximum
__$crawl [on/off]__
Activate real-time image scraping in this channel
__$crawl past__
Scraping past images
__$image [username]__
Get a random image uploaded by [username]
    `
	description := &discord.Embed{
		Title:       "Pamo_bot commands list",
		Description: strings.Replace(desc, "$", prefix, -1),
		Fields: []discord.EmbedField{
			{Name: "__Photo Archive__", Value: "https://vrc.nupa.moe"},
		},
	}

	return description, nil
}
