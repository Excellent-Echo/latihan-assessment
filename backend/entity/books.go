package entity

import "time"

type Books struct {
	ID        int       `gorm: "primaryKey" json "id"`
	UserID    []Users   `gorm: "foreignKey:ID"`
	Title     string    `json: "title"`
	Author    string    `json: "author"`
	Year      int       `json: "year"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
