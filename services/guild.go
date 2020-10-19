package services

import (
	"log"
	"strconv"

	"github.com/diamondburned/arikawa/discord"
	"github.com/nupamore/pamo_bot/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ScrapingChannelIDs : scraping channels
var ScrapingChannelIDs map[discord.ChannelID]bool

// AutoTranslateChannelIDs : auto translate channels
var AutoTranslateChannelIDs map[discord.ChannelID]bool

// InitChannelsInfo : get channels info from server
func InitChannelsInfo() {
	// auto translate channels
	AutoTranslateChannelIDs = map[discord.ChannelID]bool{
		681470820220010497: true,
		507170236265398272: true,
		662308553494626307: true,
	}

	// scraping channels
	guilds, err := models.DiscordGuilds(
		qm.Where("status = 'WATCH'"),
	).All(DB)

	if err != nil {
		log.Println("ScrapingChannesl init fail")
		panic(err)
	}

	ScrapingChannelIDs = map[discord.ChannelID]bool{}

	for _, guild := range guilds {
		idStr := *guild.ScrapChannelID.Ptr()
		id, _ := strconv.ParseUint(idStr, 10, 64)

		ScrapingChannelIDs[discord.ChannelID(id)] = true
	}
}

// AddScrapingChannel : add scraping channel
func AddScrapingChannel(guildID discord.GuildID, channelID discord.ChannelID) {
	guild, err := models.DiscordGuilds(
		qm.Where("guild_id=?", guildID),
	).One(DB)

	guild.ScrapChannelID = null.StringFrom(string(channelID))
	guild.Status = null.StringFrom("WATCH")
	guild.Update(DB, boil.Infer())

	if err != nil {
		log.Println(err)
	} else {
		ScrapingChannelIDs[channelID] = true
	}
}

// RemoveScrapingChannel : remove scraping channel
func RemoveScrapingChannel(guildID discord.GuildID, channelID discord.ChannelID) {
	guild, err := models.DiscordGuilds(
		qm.Where("guild_id=?", guildID),
	).One(DB)

	guild.ScrapChannelID = null.NewString("", false)
	guild.Status = null.StringFrom("STOP")
	guild.Update(DB, boil.Infer())

	if err != nil {
		log.Println(err)
	} else {
		delete(ScrapingChannelIDs, channelID)
	}
}
