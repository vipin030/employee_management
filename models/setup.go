package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user=interface password=interface dbname=employee sslmode=disable"
	database, err := gorm.Open("postgres", dsn)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Employee{})

	DB = database
}
