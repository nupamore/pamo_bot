package services

import (
	"log"
	"regexp"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/nupamore/pamo_bot/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
func ScrapImage(m discord.Message) error {
	file := m.Attachments[0]
	var image models.DiscordImage
	image.GuildID = null.StringFrom(string(m.GuildID))
	image.ChannelID = null.StringFrom(string(m.ChannelID))
	image.FileID = string(file.ID)
	image.FileName = null.StringFrom(file.Filename)
	image.RegDate = null.TimeFrom(time.Time(m.Timestamp))
	image.ArchiveDate = null.TimeFrom(time.Now())
	image.OwnerID = null.StringFrom(string(m.Author.ID))
	image.OwnerName = null.StringFrom(m.Author.Username)
	image.OwnerAvatar = null.StringFrom(m.Author.Avatar)

	err := image.Insert(DB, boil.Infer())

	isDuplicate, _ := regexp.MatchString("Error 1062", err.Error())
	if err != nil && !isDuplicate {
		log.Println(err)
	}

	return nil
}

// CrawlImages : scrap past images
func CrawlImages(channelID discord.ChannelID, messageID discord.MessageID) (discord.MessageID, error) {
	messages, err := DiscordAPI.MessagesBefore(channelID, messageID, 100)
	if err != nil {
		log.Println(err)
	}

	for _, m := range messages {
		if len(m.Attachments) > 0 && !m.Author.Bot {
			ScrapImage(m)
		}
	}

	if len(messages) == 0 {
		return discord.NullMessageID, err
	}
	return messages[len(messages)-1].ID, err
}
