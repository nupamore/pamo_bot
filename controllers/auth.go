package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/services"
)

// Login : [GET] /auth/login
func (ctrl *Controller) Login(c *fiber.Ctx) error {
	store := services.Sessions.Get(c)
	defer store.Save()

	url := services.GetLoginURL(store.ID())
	return c.Redirect(url)
}

// LoginCallback : [GET] /auth/callback
func (ctrl *Controller) LoginCallback(c *fiber.Ctx) error {
	store := services.Sessions.Get(c)

	if c.Query("state") != store.ID() {
		return ctrl.SendError(c, AuthError, errors.New("Invalid session state"))
	}

	token, err := services.Authenticate(c.Query("code"))
	if err != nil {
		return ctrl.SendError(c, AuthError, errors.New("OAuth login fail"))
	}

	auth := token.TokenType + " " + token.AccessToken
	store.Set("Authorization", auth)
	user, _ := services.GetUserInfo(auth)
	store.Set("UserID", user.ID)
	defer store.Save()

	return c.Redirect(configs.Env["WEB_URL"])
}

// Logout : [GET] /auth/logout
func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	store := services.Sessions.Get(c)
	store.Destroy()
	return c.JSON(Response{})
}
