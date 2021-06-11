package entity

import (
	"time"
)

type Book struct {
	ID        uint `gorm:primaryKey;autoIncrement json:"id"`
	UserID    uint
	Title     string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
