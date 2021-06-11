package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
}
