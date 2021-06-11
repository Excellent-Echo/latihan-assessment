package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	DateBirth time.Time `json:"date_birth"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"creted_at"`
	UpdateAt  time.Time `json:"update_at"`
	Books     []Books   `gorm:"foreignKey:UserID"`
}
