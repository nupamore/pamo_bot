package controllers

import "github.com/gofiber/fiber/v2"

// Middleware : Middleware
func (ctrl *Controller) Middleware(c *fiber.Ctx) error {
	return c.Next()
}
