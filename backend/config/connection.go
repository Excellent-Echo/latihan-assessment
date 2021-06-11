package config

import (
	"books-project/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "root:@tcp(localhost)/assessment_excellent-echo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&entity.Book{})
	db.AutoMigrate(&entity.User{})
	return db
}
