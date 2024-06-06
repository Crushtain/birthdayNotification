package app

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Crushtain/birthdayNotification/server/iternal/config"
	"github.com/Crushtain/birthdayNotification/server/iternal/database"
)

type App struct {
	Config *config.Config
	// Storage storage.Storage
	Database *database.Database
	Fiber    *fiber.App
}

func NewApp(config *config.Config) *App {
	return &App{
		Config:   config,
		Database: database.DBConnect(config.DatabasePath),
		Fiber:    fiber.New(),
	}
}
