package migration

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Address   string
	DateBirth time.Time
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Book      []Book `gorm:"foreignKey:UserID"`
}

type Book struct {
	ID        int `gorm:"primaryKey" json:"id"`
	Tittle    string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
}
