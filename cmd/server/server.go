package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/arl/statsviz"
	"github.com/diamondburned/arikawa/v2/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nupamore/pamo_bot/configs"
	"github.com/nupamore/pamo_bot/controllers"
	"github.com/nupamore/pamo_bot/services"
)

func router(app *fiber.App) {
	// healthcheck
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	ctrl := controllers.Controller{}
	api := app.Group("/api")
	// test
	api.Get("/links/:linkID", ctrl.GetLink)
	api.Put("/links/:linkID", ctrl.LogLink)
	// auth
	api.Get("/auth/login", ctrl.Login)
	api.Get("/auth/callback", ctrl.LoginCallback)
	api.Get("/auth/logout", ctrl.Logout)
	// auth required
	v1 := api.Group("/v1", ctrl.AuthFilter)
	v1.Get("/me", ctrl.GetMyInfo)
	v1.Get("/guilds", ctrl.GetGuilds)
	v1.Get("/guilds/:guildID", ctrl.GetGuild)
	v1.Get("/guilds/:guildID/uploaders", ctrl.GetUploaders)
	v1.Get("/guilds/:guildID/images", ctrl.GetImages)
	v1.Get("/links", ctrl.GetLinks)
	v1.Post("/links", ctrl.InitLinks)
	v1.Put("/links/:linkID", ctrl.UpdateLink)
}

func main() {
	if configs.Env["DEBUG_PORT"] != "" {
		go func() {
			statsviz.RegisterDefault()
			log.Fatal(http.ListenAndServe(configs.Env["DEBUG_PORT"], nil))
		}()
	}
	services.DBsetup()
	services.Auth.Setup()
	services.DiscordAPI = api.NewClient("Bot " + configs.Env["BOT_TOKEN"])

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     configs.Env["WEB_URL"],
		AllowCredentials: true,
	}))
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			isAuth, _ := regexp.MatchString("/api/auth", c.Path())
			notLogin := c.Cookies("session_id") == ""
			notGET := c.Method() != "GET"
			return isAuth || notLogin || notGET
		},
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprintf("%s%s%s", c.Cookies("session_id"), c.Method(), c.OriginalURL())
		},
	}))
	router(app)
	app.Listen(configs.Env["SERVER_PORT"])
}
