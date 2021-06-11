package entity

import "time"

type Books struct {
	ID int `gorm:"primaryKey" json:"id"`
	UserID int `json:"user_id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year int `json:"year"`
	CreatedAt time.Time `json:"creted_at"`
	UpdateAt time.Time `json:"update_at"`
}