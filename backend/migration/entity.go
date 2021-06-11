package migration

import "time"

type User struct {
	ID int `gorm:"primaryKey" json:"id"`
	Address string `json:"address"`
	Date_birth time.Time `json:"date_birth"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"creted_at"`
	UpdateAt time.Time `json:"update_at"`
	Books []Books `gorm:"foreignKey:UserID"`
}

type Books struct {
	ID int `gorm:"primaryKey" json:"id"`
	UserID int `json:"user_id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Year int `json:"year"`
	CreatedAt time.Time `json:"creted_at"`
	UpdateAt time.Time `json:"update_at"`
}