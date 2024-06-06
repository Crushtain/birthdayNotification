package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Crushtain/birthdayNotification/iternal/models"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}

func DBConnect(dbPath string) *Database {
	db, err := gorm.Open(postgres.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connect to database")
		return nil
	}
	database := NewDatabase(db)
	return database
}

func (d *Database) GetAllUsers() []*models.User {
	var user []*models.User
	d.DB.Find(&user)
	return user
}

func (d *Database) CreateUser(user *models.User) error {
	return d.DB.Create(user).Error
}
