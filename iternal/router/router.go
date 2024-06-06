package router

import (
	"github.com/Crushtain/birthdayNotification/iternal/app"
)

func Route(app *app.App) {

	app.Fiber.Get("/", app.MainPage)
	app.Fiber.Post("/user", app.Register)

}
