package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/Crushtain/birthdayNotification/server/iternal/models"
)

func (a *App) Register(c *fiber.Ctx) error {
	newUser := new(models.User)
	context := fiber.Map{
		"status": "OK",
		"msg":    "Add a user",
	}

	if err := c.BodyParser(&newUser); err != nil {
		log.Fatal("Error to parse the user info")
		context["status"] = "Bad request"
		context["msg"] = "Error to parse the user info"
		return c.JSON(context)
	}

	err := a.Database.CreateUser(newUser)
	if err != nil {
		context["status"] = "Bad request"
		context["msg"] = "Error to save user to database"
		return c.JSON(context)
	}

	c.Status(201)
	return c.JSON(context)
}
