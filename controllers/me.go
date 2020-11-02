package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

// GetMyInfo : [GET] /me
func (ctrl *Controller) GetMyInfo(c *fiber.Ctx) error {
	store := services.Auth.Sessions.Get(c)
	auth := store.Get("Authorization")

	user, err := services.Auth.Info(auth.(string))
	if err != nil {
		return ctrl.SendError(c, EtcError, err)
	}

	return c.JSON(Response{Data: user})
}
