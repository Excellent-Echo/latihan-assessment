package migration

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	Address   string
	BirthDate time.Time
	Books     []Book
}

type Book struct {
	gorm.Model
	UserID uint
	Title  string
	Author string
	Year   int
}
