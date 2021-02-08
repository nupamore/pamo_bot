package commands

import (
	"time"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/services"
	"github.com/nupamore/pamo_bot/utils"
)

const (
	refreshEmoji = "🔄"
	timeout      = 1 * time.Minute
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
	s := cmd.Ctx.Session
	var ownerName string
	if arg != "" {
		ownerName = string(arg)
	}

	embed := discord.Embed{}

	// query a random image
	image, err := services.Image.Random(m.GuildID, ownerName)
	if err != nil {
		embed.Description = "Couldn't find any image"
		s.SendMessage(m.ChannelID, "", &embed)
		return nil
	}

	embed = *image2embed(image)
	msg, err := s.SendMessage(m.ChannelID, "", &embed)
	s.React(msg.ChannelID, msg.ID, refreshEmoji)
	time.Sleep(time.Millisecond * 250)

	// add handler
	remove := s.AddHandler(func(r *gateway.MessageReactionAddEvent) {
		if r.MessageID != msg.ID || r.Emoji.Name != refreshEmoji {
			return
		}
		image, err := services.Image.Random(m.GuildID, ownerName)
		if err != nil {
			return
		}
		embed = *image2embed(image)
		s.EditMessage(msg.ChannelID, msg.ID, "", &embed, false)
		s.DeleteUserReaction(
			msg.ChannelID,
			msg.ID,
			r.UserID,
			refreshEmoji,
		)
	})

	// remove handler
	time.Sleep(timeout)
	s.DeleteReactions(
		msg.ChannelID,
		msg.ID,
		refreshEmoji,
	)
	defer remove()

	return err
}
