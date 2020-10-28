package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/controllers"
	"github.com/nupamore/pamo_bot/services"
)

func router(app *fiber.App) {
	ctrl := controllers.Controller{}
	// external
	app.Get("/randomImage/:guildID", ctrl.GetRandomImage)
	// auth
	app.Get("/auth/login", ctrl.Login)
	app.Get("/auth/callback", ctrl.LoginCallback)
	app.Get("/auth/logout", ctrl.Logout)
	// internal
	api := app.Group("/api/v1", ctrl.Middleware)
	api.Get("/me", ctrl.GetMyInfo)
	api.Get("/guilds", ctrl.GetGuilds)
	api.Get("/guilds/:guildID", ctrl.GetGuild)
	api.Get("/guilds/:guildID/uploaders", ctrl.GetUploaders)
	api.Get("/guilds/:guildID/images", ctrl.GetImages)
}

func main() {
	services.DBsetup()
	services.AuthSetup()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     configs.Env["WEB_URL"],
		AllowCredentials: true,
	}))
	router(app)
	app.Listen(configs.Env["SERVER_PORT"])
}
