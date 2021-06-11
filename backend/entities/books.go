package entities

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Id      	uint `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	UserID 		int  `json:"user_id"`
	Title    	string `gorm:"size:255;not null" json:"title"`
	Author		string `gorm:"size:255;not null" json:"author"`
	Year  		uint	`gorm:"not null" json:"year"`
	CreatedAt 	time.Time `gorm:"type:datetime;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time	`gorm:"type:datetime;default:current_timestamp" json:"updated_at"`
}