package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/zmcnellis/go-api-template/internal/models"
)

var db *gorm.DB

// NewDB ...
func NewDB() *gorm.DB {
	setupEnvironment()
	setupConnection()
	db.AutoMigrate(&models.Resource{})

	return db
}

func setupEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupConnection() {
	var err error
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("HOST")+" user="+os.Getenv("USER")+
			" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+os.Getenv("PASSWORD"))

	if err != nil {
		panic("failed to connect database")
	}
}
