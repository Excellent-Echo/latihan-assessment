package migration

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	DateBirth time.Time
	Address   string
	CreateAt  time.Time
	UpdateAt  time.Time
	Books     []Book
}

type Book struct {
	ID     int `gorm:"primaryKey"`
	Title  string
	Author string
	Year   int
	UserID int
}
