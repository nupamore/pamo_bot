package controllers

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

// GetLinks : [GET] /links
func (ctrl *Controller) GetLinks(c *fiber.Ctx) error {
	store := services.Auth.Sessions.Get(c)
	id := store.Get("UserID")
	ownerID, _ := discord.ParseSnowflake(id.(string))
	linkIDs, _ := services.Link.List(discord.UserID(ownerID))

	return c.JSON(Response{Data: linkIDs})
}

// InitLinks : [POST] /links
func (ctrl *Controller) InitLinks(c *fiber.Ctx) error {
	store := services.Auth.Sessions.Get(c)
	id := store.Get("UserID")
	ownerID, _ := discord.ParseSnowflake(id.(string))
	linkIDs, _ := services.Link.Create(discord.UserID(ownerID))

	return c.JSON(Response{Data: linkIDs})
}

// UpdateLink : [PUT] /links/:linkID
func (ctrl *Controller) UpdateLink(c *fiber.Ctx) error {
	store := services.Auth.Sessions.Get(c)
	id := store.Get("UserID")
	ownerID, _ := discord.ParseSnowflake(id.(string))
	linkID := c.Params("linkID")

	link, err := services.Link.Update(linkID, discord.UserID(ownerID), c.Body())
	if err != nil {
		return ctrl.SendError(c, DBError, err)
	}

	return c.JSON(Response{Data: link})
}
