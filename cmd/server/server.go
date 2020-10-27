package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/nupamore/pamo_bot/controllers"
	"github.com/nupamore/pamo_bot/services"
)

func init() {
	godotenv.Load("configs/.env")
}

func main() {
	services.DBsetup()
	services.AuthSetup()
	app := fiber.New()
	ctrl := controllers.Controller{}
	app.Use(recover.New())

	app.Get("/auth/login", ctrl.Login)
	app.Get("/auth/callback", ctrl.LoginCallback)
	app.Get("/auth/logout", ctrl.Logout)

	api := app.Group("/api/v1", ctrl.Middleware)
	api.Get("/me", ctrl.GetMyInfo)
	api.Get("/guilds", ctrl.GetGuilds)
	api.Get("/guilds/:guildID", ctrl.GetGuild)
	api.Put("/guilds/:guildID", ctrl.UpdateGuild)
	api.Get("/guilds/:guildID/uploaders", ctrl.GetUploaders)

	app.Listen(os.Getenv("WEB_PORT"))
}
