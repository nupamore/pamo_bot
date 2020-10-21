package services

import (
	"log"
	"regexp"
	"strconv"
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

	return image, err
}

// ScrapImage : save image info to server
func ScrapImage(m discord.Message) error {
	file := m.Attachments[0]
	image := models.DiscordImage{
		GuildID:     null.StringFrom(strconv.FormatUint(uint64(m.GuildID), 10)),
		ChannelID:   null.StringFrom(strconv.FormatUint(uint64(m.ChannelID), 10)),
		FileID:      string(file.ID),
		FileName:    null.StringFrom(file.Filename),
		RegDate:     null.TimeFrom(time.Time(m.Timestamp)),
		ArchiveDate: null.TimeFrom(time.Now()),
		OwnerID:     null.StringFrom(strconv.FormatUint(uint64(m.Author.ID), 10)),
		OwnerName:   null.StringFrom(m.Author.Username),
		OwnerAvatar: null.StringFrom(m.Author.Avatar),
	}

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
	if err != nil || len(messages) == 0 {
		return discord.NullMessageID, err
	}

	for _, m := range messages {
		if len(m.Attachments) > 0 && !m.Author.Bot {
			ScrapImage(m)
		}
	}

	return messages[len(messages)-1].ID, err
}

// Uploader : uploader model
type Uploader struct {
	OwnerID     string      `json:"id"`
	OwnerName   string      `json:"name"`
	OwnerAvatar null.String `json:"avatar"`
}

// GetImageUploaders : get uploaders in guild
func GetImageUploaders(guildID discord.GuildID) ([]Uploader, error) {
	uploaders := []Uploader{}
	images, err := models.DiscordImages(
		qm.Select("owner_id", "owner_name", "owner_avatar"),
		qm.Where("guild_id = ?", guildID),
		qm.GroupBy("owner_id"),
	).All(DB)

	for _, image := range images {
		uploaders = append(uploaders, Uploader{
			OwnerID:     *image.OwnerID.Ptr(),
			OwnerName:   *image.OwnerName.Ptr(),
			OwnerAvatar: image.OwnerAvatar,
		})
	}

	return uploaders, err
}
