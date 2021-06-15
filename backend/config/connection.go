package config

import (
	"book-list/migration"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database remotemysql

// Username: EeY81k6C0r

// Database name: EeY81k6C0r

// Password: LPETBac3Dx

// Server: remotemysql.com

// Port: 3306

func Connection() *gorm.DB {

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// dsn := "root:@tcp(localhost)/book_list?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
	}

	db.AutoMigrate(&migration.Users{})
	db.AutoMigrate(&migration.Books{})

	return db
}
