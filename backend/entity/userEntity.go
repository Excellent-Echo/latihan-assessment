package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	DateBirth string    `json:"date_birth"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"creted_at"`
	UpdateAt  time.Time `json:"update_at"`
	Books     []Books   `gorm:"foreignKey:UserID"`
}

type UserInput struct {
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	DateBirth string `json:"date_birth" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type UserInputUpdate struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	DateBirth string `json:"date_birth"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
