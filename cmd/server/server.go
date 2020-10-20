package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/guilds/:guildID/uploaders", ctrl.GetUploaders)

	app.Listen(os.Getenv("WEB_PORT"))
}
