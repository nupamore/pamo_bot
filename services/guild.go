package services

import (
	"log"
	"strconv"

	"github.com/diamondburned/arikawa/discord"
	"github.com/nupamore/pamo_bot/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ScrapingChannelIDs : scraping channels
var ScrapingChannelIDs map[discord.ChannelID]bool

// GetScrapingChannelIDs : get scraping channels from server
func GetScrapingChannelIDs() {
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
