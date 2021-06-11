package config

import (
	"backend/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:pastilulus@tcp(localhost)/booklist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Book{})
	return db
}
