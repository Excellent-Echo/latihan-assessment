package entities

import (
	"time"
)

type Books struct {
	Id      	uint `gorm:"primaryKey;not null;autoIncrement"`
	Title    	string `gorm:"size:255`
	Author		string `gorm:"size:255`
	Year  		uint
	CreatedAt 	time.Time
	UpdatedAt	time.Time
}