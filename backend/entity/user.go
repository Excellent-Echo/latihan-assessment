package entity

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	Address   string
	DateBirth time.Time
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books
}
