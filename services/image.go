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

// ImageService : image service
type ImageService struct{}

// Image : imag service instance
var Image = ImageService{}

// Random : get a random image
func (s *ImageService) Random(guildID discord.GuildID, ownerName string) (*models.DiscordImage, error) {
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

// Scrap : save image info to server
func (s *ImageService) Scrap(m discord.Message, guildID discord.GuildID) error {
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

// Crawl : scrap past images
func (s *ImageService) Crawl(channelID discord.ChannelID, guildID discord.GuildID, messageID discord.MessageID) (int, discord.MessageID, error) {
	messages, err := DiscordAPI.MessagesBefore(channelID, messageID, 100)
	if err != nil || len(messages) == 0 {
		return 0, discord.NullMessageID, err
	}

	count := 0
	for _, m := range messages {
		if len(m.Attachments) > 0 && !m.Author.Bot {
			if err := s.Scrap(m, guildID); err == nil {
				count = count + 1
			}
		}
	}

	return count, messages[len(messages)-1].ID, err
}

// ImageUploader : uploader model
type ImageUploader struct {
	OwnerID     string      `json:"id"`
	OwnerName   string      `json:"name"`
	OwnerAvatar null.String `json:"avatar"`
}

// Uploaders : get uploaders in guild
func (s *ImageService) Uploaders(guildID discord.GuildID) ([]ImageUploader, error) {
	uploaders := []ImageUploader{}
	images, err := models.DiscordImages(
		qm.Select("owner_id", "owner_name", "owner_avatar"),
		qm.Where("guild_id = ?", guildID),
		qm.GroupBy("owner_id"),
	).All(DB)

	for _, image := range images {
		uploaders = append(uploaders, ImageUploader{
			OwnerID:     *image.OwnerID.Ptr(),
			OwnerName:   *image.OwnerName.Ptr(),
			OwnerAvatar: image.OwnerAvatar,
		})
	}

	return uploaders, err
}

// Count : get images count
func (s *ImageService) Count(guildID discord.GuildID) (int, error) {
	count, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
	).Count(DB)

	return int(count), err
}

// List : get images with page
func (s *ImageService) List(guildID discord.GuildID, size int, page int) (models.DiscordImageSlice, error) {
	images, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
		qm.Limit(size),
		qm.Offset(size*page),
	).All(DB)

	return images, err
}
