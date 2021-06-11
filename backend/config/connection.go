package config

import (
	"book-list/migration"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "root:@tcp(localhost)/book_list"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
	}

	db.AutoMigrate(&migration.Users{})
	db.AutoMigrate(&migration.Books{})

	return db
}
