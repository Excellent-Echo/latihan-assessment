package entity

import "time"

type Books struct {
	ID        int       `gorm:"Primarykey" json:"id"`
	UserID    int       `json:"UserID"`
	Tittle    string    `json:"tittle"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookInput struct {
	Tittle string `json:"tittle"`
	Author string `json:"author"`
}
