package events

import (
	"log"
	"strconv"
	"time"

	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/services"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/diamondburned/arikawa/gateway"
)

// GuildCreated : join guild
func GuildCreated(g *gateway.GuildCreateEvent) {
	// exist already
	if _, exist := services.GuildIDs[g.ID]; exist {
		return
	}

	// not exist but comeback
	guild, _ := services.GetGuildInfo(g.ID)
	if guild != nil {
		log.Println("Comeback guild")
		guild.Status = null.StringFrom("COMEBACK")
		guild.ModDate = null.TimeFrom(time.Now())
		guild.Update(services.DB, boil.Infer())
		return
	}

	// new guild
	log.Println("New guild")
	guild = &models.DiscordGuild{
		GuildID:        strconv.FormatUint(uint64(g.ID), 10),
		GuildName:      null.StringFrom(g.Name),
		ScrapChannelID: null.NewString("", false),
		Status:         null.StringFrom("CREATED"),
		RegUser:        null.StringFrom(strconv.FormatUint(uint64(g.OwnerID), 10)),
		RegDate:        null.TimeFrom(time.Now()),
		ModUser:        null.StringFrom(strconv.FormatUint(uint64(g.OwnerID), 10)),
		ModDate:        null.TimeFrom(time.Now()),
	}
	if err := guild.Insert(services.DB, boil.Infer()); err != nil {
		services.GuildIDs[g.ID] = true
	}
}

// GuildDeleted : exit guild
func GuildDeleted(g *gateway.GuildDeleteEvent) {
	log.Println("Bot is kicked")

	guild, _ := services.GetGuildInfo(g.ID)
	guild.Status = null.StringFrom("KICKED")
	guild.ModDate = null.TimeFrom(time.Now())
	guild.Update(services.DB, boil.Infer())

	services.RemoveScrapingChannel(g.ID)
	delete(services.GuildIDs, g.ID)
}
