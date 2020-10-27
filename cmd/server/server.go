package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/controllers"
	"github.com/nupamore/pamo_bot/services"
)

func main() {
	services.DBsetup()
	services.AuthSetup()
	app := fiber.New()
	ctrl := controllers.Controller{}
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     configs.Env["WEB_URL"],
		AllowCredentials: true,
	}))

	app.Get("/auth/login", ctrl.Login)
	app.Get("/auth/callback", ctrl.LoginCallback)
	app.Get("/auth/logout", ctrl.Logout)

	api := app.Group("/api/v1", ctrl.Middleware)
	api.Get("/me", ctrl.GetMyInfo)
	api.Get("/guilds", ctrl.GetGuilds)
	api.Get("/guilds/:guildID", ctrl.GetGuild)
	api.Put("/guilds/:guildID", ctrl.UpdateGuild)
	api.Get("/guilds/:guildID/uploaders", ctrl.GetUploaders)

	app.Listen(configs.Env["SERVER_PORT"])
}
