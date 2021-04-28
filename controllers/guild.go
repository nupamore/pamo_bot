package controllers

import (
	"strconv"
	"sync"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/models"
	"github.com/nupamore/pamo_bot/services"
)

// GetGuilds : [GET] /guilds
func (ctrl *Controller) GetGuilds(c *fiber.Ctx) error {
	sess, _ := services.Auth.Store.Get(c)
	auth := sess.Get("Authorization")

	var err error
	oauthGuilds := []*services.DiscordGuild{}
	serverGuilds := []*models.DiscordGuild{}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		oauthGuilds, err = services.Auth.Guilds(auth.(string))
	}()
	go func() {
		defer wg.Done()
		serverGuilds, err = services.Guild.All()
	}()
	wg.Wait()

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

	var all int
	var pageMeta PageMeta
	var images models.DiscordImageSlice

	var wg sync.WaitGroup
	wg.Add(2)
	// images count
	go func() {
		defer wg.Done()
		all, err = services.Image.Count(discord.GuildID(guildID))
		if all < size*page {
			page = all / size
		}
		pageMeta = PageMeta{
			Size: size,
			Page: page,
			All:  all,
		}
	}()
	// image list
	go func() {
		defer wg.Done()
		images, err = services.Image.List(discord.GuildID(guildID), owner, size, page)
	}()
	wg.Wait()

	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{
		PageMeta: &pageMeta,
		Data:     images,
	})
}
