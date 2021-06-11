package entity

import "time"

type Users struct {
	ID        int    `gorm:"Primarykey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique"`
	DateBirth time.Time
	Address   string `json:"address"`
}
