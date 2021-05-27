package services

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"  // gif
	_ "image/jpeg" // jpg
	_ "image/png"  // png
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/buckket/go-blurhash"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/utils"
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
		qm.And("status IS NULL"),
		qm.OrderBy("rand()"),
	).One(DB)

	return image, err
}

// Scrap : save image info to server
func (s *ImageService) Scrap(m discord.Message, guildID discord.GuildID) error {
	file := m.Attachments[0]
	fileID := strconv.FormatUint(uint64(file.ID), 10)
	image, err := models.FindDiscordImage(DB, fileID)

	if image != nil {
		return errors.New("duplicate")
	}

	blur, _ := BlurHash(utils.DiscordMediaServer(
		strconv.FormatUint(uint64(m.ChannelID), 10),
		strconv.FormatUint(uint64(file.ID), 10),
		file.Filename,
		"width=48&height=27",
	))
	image = &models.DiscordImage{
		FileID:      fileID,
		FileName:    null.StringFrom(file.Filename),
		OwnerName:   null.StringFrom(m.Author.Username),
		OwnerID:     null.StringFrom(strconv.FormatUint(uint64(m.Author.ID), 10)),
		OwnerAvatar: null.StringFrom(m.Author.Avatar),
		GuildID:     null.StringFrom(strconv.FormatUint(uint64(guildID), 10)),
		ChannelID:   null.StringFrom(strconv.FormatUint(uint64(m.ChannelID), 10)),
		Width:       null.StringFrom(strconv.FormatUint(uint64(file.Width), 10)),
		Height:      null.StringFrom(strconv.FormatUint(uint64(file.Height), 10)),
		Blurhash:    null.StringFromPtr(blur),
		RegDate:     null.TimeFrom(time.Time(m.Timestamp)),
		ArchiveDate: null.TimeFrom(time.Now()),
	}
	// err = image.Upsert(DB, boil.Whitelist("blurhash"), boil.Infer())
	err = image.Insert(DB, boil.Infer())

	return err
}

// Crawl : scrap past images
func (s *ImageService) Crawl(channelID discord.ChannelID, guildID discord.GuildID, messageID discord.MessageID) (int, discord.MessageID, error) {
	messages, err := DiscordAPI.MessagesBefore(channelID, messageID, 100)
	if err != nil || len(messages) == 0 {
		return 0, discord.NullMessageID, err
	}

	var wg sync.WaitGroup
	var count uint32 = 0

	for _, m := range messages {
		if len(m.Attachments) > 0 && !m.Author.Bot {
			wg.Add(1)
			go func(m discord.Message) {
				defer wg.Done()
				if err := s.Scrap(m, guildID); err == nil {
					atomic.AddUint32(&count, 1)
				}
			}(m)
		}
	}
	wg.Wait()

	return int(atomic.LoadUint32(&count)), messages[len(messages)-1].ID, err
}

// Uploader : uploader model
type Uploader struct {
	OwnerID     string      `json:"id"`
	OwnerName   string      `json:"name"`
	OwnerAvatar null.String `json:"avatar"`
}

// Uploaders : get uploaders in guild
func (s *ImageService) Uploaders(guildID discord.GuildID) ([]Uploader, error) {
	uploaders := []Uploader{}
	images, err := models.DiscordImages(
		qm.Select("owner_id", "owner_name", "owner_avatar"),
		qm.Where("guild_id = ?", guildID),
		qm.GroupBy("owner_id"),
		qm.OrderBy("owner_name"),
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

// Count : get images count
func (s *ImageService) Count(guildID discord.GuildID) (int, error) {
	count, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
		qm.And("status IS NULL"),
	).Count(DB)

	return int(count), err
}

// List : get images with page
func (s *ImageService) List(guildID discord.GuildID, ownerName string, size int, page int) (models.DiscordImageSlice, error) {
	images, err := models.DiscordImages(
		qm.Where("guild_id = ?", guildID),
		qm.And("owner_name LIKE ?", fmt.Sprintf("%%%s%%", ownerName)),
		qm.And("status IS NULL"),
		qm.Limit(size),
		qm.Offset(size*page),
		qm.OrderBy("reg_date DESC"),
	).All(DB)

	return images, err
}

// Delete : delete images
func (s *ImageService) Delete(ownerID discord.UserID, guildID discord.GuildID, fileIDs []interface{}) (int, error) {
	images, err := models.DiscordImages(
		qm.Where("owner_id = ? ", ownerID),
		qm.And("guild_id = ?", guildID),
		qm.AndIn("file_id IN ?", fileIDs...),
	).All(DB)

	count, err := images.UpdateAll(DB, models.M{"status": "DELETED"})

	return int(count), err
}

// DeleteMaster : delete images master permission
func (s *ImageService) DeleteMaster(guildID discord.GuildID, fileIDs []interface{}) (int, error) {
	images, err := models.DiscordImages(
		qm.And("guild_id = ?", guildID),
		qm.AndIn("file_id IN ?", fileIDs...),
	).All(DB)

	count, err := images.UpdateAll(DB, models.M{"status": "DELETED"})

	return int(count), err
}

// BlurHash : url to blurhash
func BlurHash(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Not found")
	}
	if match, _ := regexp.MatchString("image", resp.Header.Get("Content-type")); !match {
		return nil, errors.New("Not image")
	}

	// image
	loadedImage, _, err := image.Decode(resp.Body)
	blur, err := blurhash.Encode(4, 3, loadedImage)
	if err != nil {
		return nil, err
	}

	return &blur, nil
}
