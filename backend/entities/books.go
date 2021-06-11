package entities

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Id      	uint `gorm:"primaryKey;not null;autoIncrement"`
	UserId 		uint 
	Title    	string `gorm:"size:255`
	Author		string `gorm:"size:255`
	Year  		uint
	CreatedAt 	time.Time
	UpdatedAt	time.Time
	UserID int
}