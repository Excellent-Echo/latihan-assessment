package entities

import (
	"time"
)

type Users struct {
	Id      uint   `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	Name    string `gorm:"size:255;not null" json:"name"`
	Address string `gorm:"size:255;not null" json:"address"`
	DateBirth time.Time `gorm:"size:255;not null" json:"date_birth"`
	Email     string `gorm:"unique;size:255:not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `gorm:"type:datetime;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time	`gorm:"type:datetime;default:current_timestamp" json:"updated_at"`
	Book []Books `gorm:"foreignKey:UserID" json:"book"`
}