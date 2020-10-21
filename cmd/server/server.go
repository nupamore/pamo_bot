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
	app := fiber.New()
	ctrl := controllers.Controller{}
	app.Use(recover.New())

	api := app.Group("/api/v1", ctrl.Middleware)
	api.Get("/guilds", ctrl.GetGuilds)
	api.Get("/guilds/:guildID", ctrl.GetGuild)
	api.Put("/guilds/:guildID", ctrl.UpdateGuild)
	api.Get("/guilds/:guildID/uploaders", ctrl.GetUploaders)

	app.Listen(os.Getenv("WEB_PORT"))
}
