package migration

import "time"

type Users struct {
	ID        int `gorm:"Primarykey"`
	Name      string
	Address   string
	DateBirth time.Time
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []Books `gorm:"Foreignkey:UserID"`
}

type Books struct {
	ID        int `gorm:"Primarykey"`
	UserID    int
	Tittle    string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
