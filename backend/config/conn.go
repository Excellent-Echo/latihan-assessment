package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
)

func ConnDB() *gorm.DB  {

	//dsn (Data Source Name)
	dsn := "root:@tcp(localhost)/book_list_excellent_echo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	fmt.Println("database connected")

	return db

}