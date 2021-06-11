package entity

import "time"

type Users struct {
	ID        int       `gorm:"Primarykey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	DateBirth time.Time `json:"date_birth"`
	Address   string    `json:"address"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
