package discord

import (
	"fmt"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Image : get a random image from this guild
func (cmd *Commands) Image(m *gateway.MessageCreateEvent, arg bot.RawArguments) (*discord.Embed, error) {
	var ownerName string
	if arg != "" {
		ownerName = string(arg)
	}

	embed := discord.Embed{}

	// query a random image
	image, err := Service.GetRandomImage(m.GuildID, ownerName)
	if err != nil {
		embed.Description = "Couldn't find any image"
		return &embed, nil
	}

	embed.Footer = &discord.EmbedFooter{Text: "📷 " + *image.OwnerName.Ptr()}
	embed.Timestamp = discord.NewTimestamp(*image.RegDate.Ptr())
	embed.Image = &discord.EmbedImage{URL: fmt.Sprintf(
		"https://cdn.discordapp.com/attachments/%s/%s/%s",
		*image.ChannelID.Ptr(),
		image.FileID,
		*image.FileName.Ptr(),
	)}

	return &embed, err
}
