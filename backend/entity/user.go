package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Books     []Book
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `gorm:"index" json:"-"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserInput struct {
	Name      string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required"`
	DateBirth time.Time `json:"date_birth" binding:"required"`
}

type UpdateUserInput struct {
	Name      string    `json:"name"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `json:"email"`
}
