package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/services"
)

// AuthFilter : api middleware
func (ctrl *Controller) AuthFilter(c *fiber.Ctx) error {
	// oauth check
	sess, _ := services.Auth.Store.Get(c)
	auth := sess.Get("Authorization")
	if auth == nil {
		return ctrl.SendError(c, AuthError, errors.New("Need login"))
	}

	return c.Next()
}
