package config

import (
	"github.com/marwanjuna/latihan-assessment/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:marwanajunolii@tcp(localhost)/booklistapp"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Book{})

	return db
}
