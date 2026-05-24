package configs

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("exam.db"), &gorm.Config{})

	if err != nil {
		//log.Fatal("failed to connect database")
		log.Fatal(err)
	}

	DB = database

	log.Println("database connected")
}