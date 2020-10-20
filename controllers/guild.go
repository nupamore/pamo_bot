package controllers

import (
	"encoding/json"

	"github.com/diamondburned/arikawa/discord"
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

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

	res := Response{
		Code: 0,
		Data: uploaders,
	}
	r, _ := json.Marshal(res)
	return c.Send(r)
}
