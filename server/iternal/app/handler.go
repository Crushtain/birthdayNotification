package app

import (
	"github.com/gofiber/fiber/v2"
)

func (a *App) MainPage(c *fiber.Ctx) error {
	response := fiber.Map{
		"statusText": "ok",
		"msg":        "Add a birthday",
	}
	users := a.Database.GetAllUsers()
	response["users"] = users
	c.Status(201)
	return c.JSON(response)

}
