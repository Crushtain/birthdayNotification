package router

import (
	"github.com/Crushtain/birthdayNotification/server/iternal/app"
)

func Route(app *app.App) {

	app.Fiber.Get("/", app.MainPage)
	app.Fiber.Get("/:id", app.Find)
	app.Fiber.Post("/login", app.Register)

}
