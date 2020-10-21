package controllers

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

// GetGuilds : [GET] /guilds
func (ctrl *Controller) GetGuilds(c *fiber.Ctx) error {
	guilds, err := services.GetAllGuildsInfo()
	if err != nil {
		return ctrl.SendError(c, DBError, err)
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
