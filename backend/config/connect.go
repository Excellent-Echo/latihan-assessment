package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"latihan-assessment/backend/entity"
)

func Conn() *gorm.DB {
	dsn := "root:@tcp(localhost)/book_assesment?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.Users{})
	db.AutoMigrate(&entity.Books{})

	return db
}
