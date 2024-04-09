package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConection() {
	// db connection
	godotenv.Load()
	DNS := os.Getenv("DSN")

	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		// panic("Failed to connect to database!")
	} else {
		log.Println("Connection to database established.")
	}
}
