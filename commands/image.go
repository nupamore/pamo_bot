package commands

import (
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/services"
	"github.com/nupamore/pamo_bot/utils"
)

// Image : get a random image from this guild
func (cmd *Commands) Image(m *gateway.MessageCreateEvent, arg bot.RawArguments) (*discord.Embed, error) {
	var ownerName string
	if arg != "" {
		ownerName = string(arg)
	}

	embed := discord.Embed{}

	// query a random image
	image, err := services.Image.Random(m.GuildID, ownerName)
	if err != nil {
		embed.Description = "Couldn't find any image"
		return &embed, nil
	}

	embed.Footer = &discord.EmbedFooter{Text: "📷 " + *image.OwnerName.Ptr()}
	embed.Timestamp = discord.NewTimestamp(*image.RegDate.Ptr())
	embed.Image = &discord.EmbedImage{URL: utils.DiscordImageCDN(
		*image.ChannelID.Ptr(),
		image.FileID,
		*image.FileName.Ptr(),
	)}

	return &embed, nil
}
