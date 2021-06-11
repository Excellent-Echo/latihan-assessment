package entities

import (
	"time"
)

type User struct {
	Id      uint   `gorm:"primaryKey;not null;autoIncrement"`
	Name    string `gorm:"size:255`
	Address string `gorm:"size:255`
	DateBirth time.Time
	Email     string `gorm:"unique;size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Book []Books `gorm:foreignKey:UserID`
}