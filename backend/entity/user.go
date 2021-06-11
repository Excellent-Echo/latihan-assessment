package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Book      []Book    `json:"book"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInput struct {
	Name      string    `json:"name" binding:"required"`
	Address   string    `json:"address" binding:"required"`
	DateBirth time.Time `json:"date_birth" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required"`
}
