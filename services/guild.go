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

// GuildIDs : guilds
var GuildIDs map[discord.GuildID]bool

// ScrapingChannelIDs : scraping channels
var ScrapingChannelIDs map[discord.ChannelID]bool

// AutoTranslateChannelIDs : auto translate channels
var AutoTranslateChannelIDs map[discord.ChannelID]bool

// InitGuildsInfo : get guilds info from server
func InitGuildsInfo() {
	// auto translate channels
	AutoTranslateChannelIDs = map[discord.ChannelID]bool{
		681470820220010497: true,
		507170236265398272: true,
		662308553494626307: true,
	}

	guilds, err := GetAllGuildsInfo()

	if err != nil {
		log.Println("Guilds init fail")
		panic(err)
	}

	GuildIDs = map[discord.GuildID]bool{}
	ScrapingChannelIDs = map[discord.ChannelID]bool{}

	for _, guild := range guilds {
		guildID, _ := strconv.ParseUint(guild.GuildID, 10, 64)
		GuildIDs[discord.GuildID(guildID)] = true

		hasChannel := guild.ScrapChannelID.Valid
		if hasChannel {
			channelID, _ := strconv.ParseUint(*guild.ScrapChannelID.Ptr(), 10, 64)
			ScrapingChannelIDs[discord.ChannelID(channelID)] = true
		}
	}

	log.Printf("Guilds count: %d\n", len(GuildIDs))
	log.Printf("Crawling channels count: %d\n", len(ScrapingChannelIDs))
}

// AddScrapingChannel : add scraping channel
func AddScrapingChannel(guildID discord.GuildID, channelID discord.ChannelID) {
	guild, err := GetGuildInfo(guildID)
	guild.ScrapChannelID = null.StringFrom(strconv.FormatUint(uint64(channelID), 10))
	guild.Status = null.StringFrom("WATCH")
	guild.ModDate = null.TimeFrom(time.Now())
	guild.Update(DB, boil.Infer())

	if err != nil {
		log.Println(err)
	} else {
		ScrapingChannelIDs[channelID] = true
	}
}

// RemoveScrapingChannel : remove scraping channel
func RemoveScrapingChannel(guildID discord.GuildID) {
	guild, err := GetGuildInfo(guildID)
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
		delete(ScrapingChannelIDs, discord.ChannelID(channelID))
	}
}

// GetAllGuildsInfo : get all guilds info
func GetAllGuildsInfo() ([]*models.DiscordGuild, error) {
	guilds, err := models.DiscordGuilds(
		qm.Where("status!=?", "KICKED"),
	).All(DB)
	return guilds, err
}

// GetGuildInfo : get a guild info
func GetGuildInfo(guildID discord.GuildID) (*models.DiscordGuild, error) {
	guild, err := models.DiscordGuilds(
		qm.Where("guild_id=?", guildID),
	).One(DB)
	return guild, err
}

// UpdateGuildInfo : update a guild info
func UpdateGuildInfo(guildID discord.GuildID, options []byte) error {
	guild, err := GetGuildInfo(guildID)
	json.Unmarshal(options, &guild)
	guild.Update(DB, boil.Infer())
	return err
}
