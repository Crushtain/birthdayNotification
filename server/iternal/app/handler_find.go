package app

import "github.com/gofiber/fiber/v2"

func (a *App) Find(c *fiber.Ctx) error {
	response := fiber.Map{
		"statusText": "ok",
		"msg":        "Add a birthday",
	}

	id := c.Params("id")

	user := a.Database.GetUserByID(id)
	response["user"] = user
	c.Status(201)
	return c.JSON(response)

}
