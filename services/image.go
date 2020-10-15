package services

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/nupamore/pamo_bot/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetRandomImage : get a random image
func GetRandomImage(guildID discord.GuildID, ownerName string) (*models.DiscordImage, error) {
	if ownerName == "" {
		ownerName = "%"
	}
	image, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
		qm.And("owner_name LIKE ?", ownerName),
		qm.OrderBy("rand()"),
	).One(DB)

	if err != nil {
		return nil, err
	}

	return image, nil
}

// ScrapImage : save image info to server
func ScrapImage(m *gateway.MessageCreateEvent) error {
	file := m.Attachments[0]
	var image models.DiscordImage

	return nil
}
