package commands

import (
	"strconv"
	"time"

	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/dgwidgets"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/services"
	"github.com/nupamore/pamo_bot/utils"
)

func image2embed(image *models.DiscordImage) *discord.Embed {
	return &discord.Embed{
		Footer:    &discord.EmbedFooter{Text: "📷 " + *image.OwnerName.Ptr()},
		Timestamp: discord.NewTimestamp(*image.RegDate.Ptr()),
		Image: &discord.EmbedImage{URL: utils.DiscordImageCDN(
			*image.ChannelID.Ptr(),
			image.FileID,
			*image.FileName.Ptr(),
		)},
	}
}

// Image : get a random image from this guild
func (cmd *Commands) Image(m *gateway.MessageCreateEvent, arg bot.RawArguments) error {
	var ownerName string
	if arg != "" {
		ownerName = string(arg)
	}

	p := &dgwidgets.Paginator{
		Session:                 cmd.Ctx.Session,
		Index:                   0,
		Loop:                    false,
		DeleteMessageWhenDone:   false,
		DeleteReactionsWhenDone: false,
		ColourWhenDone:          0xFF0000,
		Widget:                  dgwidgets.NewWidget(cmd.Ctx.Session, m.ChannelID, nil),
	}
	p.Widget.Timeout = 5 * time.Minute

	embed := discord.Embed{}

	// query a random image
	image, err := services.Image.Random(m.GuildID, ownerName)
	if err != nil {
		embed.Description = "Couldn't find any image"
		p.Add(embed)
		return p.Spawn()
	}
	embed = *image2embed(image)

	p.Add(embed)

	p.Widget.Handle(dgwidgets.NavLeft, func(r *gateway.MessageReactionAddEvent) {
		if err := p.PreviousPage(); err == nil {
			p.Update()
		}
	})
	// new image
	p.Widget.Handle(dgwidgets.NavRight, func(r *gateway.MessageReactionAddEvent) {
		if err := p.NextPage(); err != nil {
			newImage, err := services.Image.Random(m.GuildID, ownerName)
			if err != nil {
				return
			}
			newEmbed := *image2embed(newImage)
			newEmbed.Description = strconv.Itoa(p.Index + 1)
			p.Add(newEmbed)
			p.NextPage()
		}
		p.Update()
	})

	return p.Spawn()
}
