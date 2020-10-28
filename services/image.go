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
func ScrapImage(m discord.Message, guildID discord.GuildID) error {
	file := m.Attachments[0]
	image := models.DiscordImage{
		FileID:      strconv.FormatUint(uint64(file.ID), 10),
		FileName:    null.StringFrom(file.Filename),
		OwnerName:   null.StringFrom(m.Author.Username),
		OwnerID:     null.StringFrom(strconv.FormatUint(uint64(m.Author.ID), 10)),
		OwnerAvatar: null.StringFrom(m.Author.Avatar),
		GuildID:     null.StringFrom(strconv.FormatUint(uint64(guildID), 10)),
		ChannelID:   null.StringFrom(strconv.FormatUint(uint64(m.ChannelID), 10)),
		Width:       null.StringFrom(strconv.FormatUint(uint64(file.Width), 10)),
		Height:      null.StringFrom(strconv.FormatUint(uint64(file.Height), 10)),
		RegDate:     null.TimeFrom(time.Time(m.Timestamp)),
		ArchiveDate: null.TimeFrom(time.Now()),
	}

	err := image.Insert(DB, boil.Infer())

	if err != nil {
		isDuplicate, _ := regexp.MatchString("Error 1062", err.Error())
		if !isDuplicate {
			log.Println(err)
		}
	}

	return err
}

// CrawlImages : scrap past images
func CrawlImages(channelID discord.ChannelID, guildID discord.GuildID, messageID discord.MessageID) (int, discord.MessageID, error) {
	messages, err := DiscordAPI.MessagesBefore(channelID, messageID, 100)
	if err != nil || len(messages) == 0 {
		return 0, discord.NullMessageID, err
	}

	count := 0
	for _, m := range messages {
		if len(m.Attachments) > 0 && !m.Author.Bot {
			if err := ScrapImage(m, guildID); err == nil {
				count = count + 1
			}
		}
	}

	return count, messages[len(messages)-1].ID, err
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

// GetImagesCount : get images count
func GetImagesCount(guildID discord.GuildID) (int, error) {
	count, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
	).Count(DB)

	return int(count), err
}

// GetImages : get images with page
func GetImages(guildID discord.GuildID, size int, page int) (models.DiscordImageSlice, error) {
	images, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
		qm.Limit(size),
		qm.Offset(size*page),
	).All(DB)

	return images, err
}
