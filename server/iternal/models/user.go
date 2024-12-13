package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       uint        `json:"id" gorm:"primaryKey; column:id"`
	Username string      `json:"username" gorm:"not null; column:username"`
	Email    string      `json:"email" gorm:"not null; column:email"`
	Birthday pgtype.Date `json:"birthday" gorm:"not null; column:bday"`
	// Subscriptions []uint      `json:"subscriptions" gorm:"column:subscriptions"`
}
