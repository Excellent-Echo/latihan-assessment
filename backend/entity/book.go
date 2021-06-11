package entity

import "time"

type Book struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Tittle    string    `json:"tittle"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
}

type BookInput struct {
	Tittle string `json:"tittle" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   string `json:"year" binding:"required"`
}

type UpdateBookInput struct {
	Tittle string `json:"tittle"`
	Author string `json:"author"`
	Year   string `json:"year"`
}
