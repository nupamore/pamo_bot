package controllers

import (
	"strconv"

	"github.com/diamondburned/arikawa/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/services"
)

// GetGuilds : [GET] /guilds
func (ctrl *Controller) GetGuilds(c *fiber.Ctx) error {
	store := services.Sessions.Get(c)
	auth := store.Get("Authorization")

	oauthGuilds, err := services.GetUsersGuilds(auth.(string))
	serverGuilds, err := services.GetAllGuildsInfo()
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	guilds := []*models.DiscordGuild{}

	for _, og := range oauthGuilds {
		for _, sg := range serverGuilds {
			if og.ID == sg.GuildID {
				guilds = append(guilds, sg)
			}
		}
	}

	return c.JSON(Response{Data: guilds})
}

// GetGuild : [GET] /guilds/:guildID
func (ctrl *Controller) GetGuild(c *fiber.Ctx) error {
	guildID, err := discord.ParseSnowflake(c.Params("guildID"))
	if err != nil {
		return ctrl.SendError(c, InvalidParamError, err)
	}
	guild, err := services.GetGuildInfo(discord.GuildID(guildID))
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{Data: guild})
}

// UpdateGuild : [PUT] /guilds/:guildID
func (ctrl *Controller) UpdateGuild(c *fiber.Ctx) error {
	guildID, err := discord.ParseSnowflake(c.Params("guildID"))
	if err != nil {
		return ctrl.SendError(c, InvalidParamError, err)
	}
	err = services.UpdateGuildInfo(discord.GuildID(guildID), c.Body())
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}
	return c.JSON(Response{})
}

// GetUploaders : [GET] /guilds/:guildID/uploaders
func (ctrl *Controller) GetUploaders(c *fiber.Ctx) error {
	guildID, err := discord.ParseSnowflake(c.Params("guildID"))
	if err != nil {
		return ctrl.SendError(c, InvalidParamError, err)
	}

	uploaders, err := services.GetImageUploaders(discord.GuildID(guildID))
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{Data: uploaders})
}

// GetImages : [GET] /guilds/:guildID/images
func (ctrl *Controller) GetImages(c *fiber.Ctx) error {
	guildID, err := discord.ParseSnowflake(c.Params("guildID"))
	size, err := strconv.Atoi(c.Query("size"))
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return ctrl.SendError(c, InvalidParamError, err)
	}

	all, err := services.GetImagesCount(discord.GuildID(guildID))
	if all < size*page {
		page = all / size
	}
	pageMeta := PageMeta{
		Size: size,
		Page: page,
		All:  all,
	}

	images, err := services.GetImages(discord.GuildID(guildID), size, page)
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{
		PageMeta: &pageMeta,
		Data:     images,
	})
}
