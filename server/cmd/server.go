package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/Crushtain/birthdayNotification/server/iternal/app"
	"github.com/Crushtain/birthdayNotification/server/iternal/config"
	"github.com/Crushtain/birthdayNotification/server/iternal/router"
)

func main() {

	cfg := config.NewConfig()

	App := app.NewApp(cfg)
	App.Fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.Route(App)

	App.Fiber.Listen(cfg.Host)
}
