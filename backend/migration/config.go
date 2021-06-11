package migration

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
	UserID int    `json:"user_id"`
}
