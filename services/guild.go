package services

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/nupamore/pamo_bot/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GuildService : guild service
type GuildService struct {
	GuildIDs                map[discord.GuildID]bool
	ScrapingChannelIDs      map[discord.ChannelID]bool
	AutoTranslateChannelIDs map[discord.ChannelID]bool
}

// Guild : guild service instance
var Guild = GuildService{}

// BotStart : get guilds info from server
func (s *GuildService) BotStart() {
	// auto translate channels
	s.AutoTranslateChannelIDs = map[discord.ChannelID]bool{
		681470820220010497: true,
		507170236265398272: true,
		662308553494626307: true,
	}

	guilds, err := s.All()

	if err != nil {
		log.Println("Guilds init fail")
		panic(err)
	}

	s.GuildIDs = map[discord.GuildID]bool{}
	s.ScrapingChannelIDs = map[discord.ChannelID]bool{}

	for _, guild := range guilds {
		guildID, _ := strconv.ParseUint(guild.GuildID, 10, 64)
		s.GuildIDs[discord.GuildID(guildID)] = true

		hasChannel := guild.ScrapChannelID.Valid
		if hasChannel {
			channelID, _ := strconv.ParseUint(*guild.ScrapChannelID.Ptr(), 10, 64)
			s.ScrapingChannelIDs[discord.ChannelID(channelID)] = true
		}
	}

	log.Printf("Guilds count: %d\n", len(s.GuildIDs))
	log.Printf("Crawling channels count: %d\n", len(s.ScrapingChannelIDs))
}

// AddScrapingChannel : add scraping channel
func (s *GuildService) AddScrapingChannel(guildID discord.GuildID, channelID discord.ChannelID) {
	guild, err := s.Info(guildID)
	guild.ScrapChannelID = null.StringFrom(strconv.FormatUint(uint64(channelID), 10))
	guild.Status = null.StringFrom("WATCH")
	guild.ModDate = null.TimeFrom(time.Now())
	guild.Update(DB, boil.Infer())

	if err != nil {
		log.Println(err)
	} else {
		s.ScrapingChannelIDs[channelID] = true
	}
}

// RemoveScrapingChannel : remove scraping channel
func (s *GuildService) RemoveScrapingChannel(guildID discord.GuildID) {
	guild, err := s.Info(guildID)
	if !guild.ScrapChannelID.Valid {
		return
	}
	channelID, _ := strconv.ParseUint(*guild.ScrapChannelID.Ptr(), 10, 64)
	guild.ScrapChannelID = null.NewString("", false)
	guild.Status = null.StringFrom("STOP")
	guild.ModDate = null.TimeFrom(time.Now())
	guild.Update(DB, boil.Infer())

	if err != nil {
		log.Println(err)
	} else {
		delete(s.ScrapingChannelIDs, discord.ChannelID(channelID))
	}
}

// All : get all guilds info
func (s *GuildService) All() ([]*models.DiscordGuild, error) {
	guilds, err := models.DiscordGuilds(
		qm.Where("status!=?", "KICKED"),
		qm.OrderBy("guild_name"),
	).All(DB)
	return guilds, err
}

// Info : get a guild info
func (s *GuildService) Info(guildID discord.GuildID) (*models.DiscordGuild, error) {
	guild, err := models.DiscordGuilds(
		qm.Where("guild_id=?", guildID),
	).One(DB)
	return guild, err
}

// Update : update a guild info
func (s *GuildService) Update(guildID discord.GuildID, options []byte) error {
	guild, err := s.Info(guildID)
	json.Unmarshal(options, &guild)
	guild.Update(DB, boil.Infer())
	return err
}
