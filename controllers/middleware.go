package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

// Middleware : Middleware
func (ctrl *Controller) Middleware(c *fiber.Ctx) error {
	// oauth check
	store := services.Sessions.Get(c)
	auth := store.Get("Authorization")
	if auth == nil {
		return ctrl.SendError(c, AuthError, errors.New("Need login"))
	}

	return c.Next()
}
