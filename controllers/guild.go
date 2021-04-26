package controllers

import (
	"strconv"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/services"
)

// GetGuilds : [GET] /guilds
func (ctrl *Controller) GetGuilds(c *fiber.Ctx) error {
	sess, _ := services.Auth.Store.Get(c)
	auth := sess.Get("Authorization")

	oauthGuilds, err := services.Auth.Guilds(auth.(string))
	serverGuilds, err := services.Guild.All()
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
	guild, err := services.Guild.Info(discord.GuildID(guildID))
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
	err = services.Guild.Update(discord.GuildID(guildID), c.Body())
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

	uploaders, err := services.Image.Uploaders(discord.GuildID(guildID))
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{Data: uploaders})
}

// GetImages : [GET] /guilds/:guildID/images
func (ctrl *Controller) GetImages(c *fiber.Ctx) error {
	guildID, err := discord.ParseSnowflake(c.Params("guildID"))
	owner := c.Query("owner")
	size, err := strconv.Atoi(c.Query("size"))
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return ctrl.SendError(c, InvalidParamError, err)
	}

	all, err := services.Image.Count(discord.GuildID(guildID))
	if all < size*page {
		page = all / size
	}
	pageMeta := PageMeta{
		Size: size,
		Page: page,
		All:  all,
	}

	images, err := services.Image.List(discord.GuildID(guildID), owner, size, page)
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{
		PageMeta: &pageMeta,
		Data:     images,
	})
}
