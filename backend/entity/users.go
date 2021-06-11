package entity

import "time"

type Users struct {
	ID        int       `gorm:"primaryKey" json: "id"`
	Name      string    `json: "name"`
	Address   string    `json: "address"`
	DateBirth time.Time `json: "date_birth"`
	Email     string    `json: "email"`
	Password  string    `json: "-"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
