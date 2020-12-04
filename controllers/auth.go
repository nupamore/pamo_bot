package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/services"
)

// Login : [GET] /auth/login
func (ctrl *Controller) Login(c *fiber.Ctx) error {
	sess, _ := services.Auth.Store.Get(c)
	sess.Save()
	url := services.Auth.LoginURL(sess.ID())
	return c.Redirect(url)
}

// LoginCallback : [GET] /auth/callback
func (ctrl *Controller) LoginCallback(c *fiber.Ctx) error {
	sess, _ := services.Auth.Store.Get(c)

	if c.Query("state") != sess.ID() {
		return ctrl.SendError(c, AuthError, errors.New("Invalid session state"))
	}

	token, err := services.Auth.Authenticate(c.Query("code"))
	if err != nil {
		return ctrl.SendError(c, AuthError, errors.New("OAuth login fail"))
	}

	auth := token.TokenType + " " + token.AccessToken
	sess.Set("Authorization", auth)
	user, _ := services.Auth.Info(auth)
	sess.Set("UserID", user.ID)
	defer sess.Save()

	return c.Redirect(configs.Env["WEB_URL"])
}

// Logout : [GET] /auth/logout
func (ctrl *Controller) Logout(c *fiber.Ctx) error {
	sess, _ := services.Auth.Store.Get(c)
	sess.Destroy()
	return c.JSON(Response{})
}
