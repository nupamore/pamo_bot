package controllers

import (
	"fmt"

	"github.com/diamondburned/arikawa/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

// GetRandomImage : [GET] /randomImage/:guildID
func (ctrl *Controller) GetRandomImage(c *fiber.Ctx) error {
	guildID, err := discord.ParseSnowflake(c.Params("guildID"))
	if err != nil {
		return ctrl.SendError(c, InvalidParamError, err)
	}

	image, err := services.Image.Random(discord.GuildID(guildID), c.Query("uploader"))
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.Redirect(fmt.Sprintf(
		"https://cdn.discordapp.com/attachments/%s/%s/%s",
		*image.ChannelID.Ptr(),
		image.FileID,
		*image.FileName.Ptr(),
	))
}

// GetLink : [GET] /links/:linkID
func (ctrl *Controller) GetLink(c *fiber.Ctx) error {
	linkID := c.Params("linkID")

	link, err := services.Link.Info(linkID)
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	ctrl.LogLink(c)

	return c.Redirect(*link.Target.Ptr())
}

// LogLink : [PUT] /links/:linkID
func (ctrl *Controller) LogLink(c *fiber.Ctx) error {
	linkID := c.Params("linkID")
	err := services.Link.Log(linkID)
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{})
}
