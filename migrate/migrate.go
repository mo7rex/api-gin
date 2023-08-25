package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mo7rex/api-gin/database"
	"github.com/mo7rex/api-gin/model"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.ConnectToDatabase()
}

func main() {
	database.DB.AutoMigrate(&model.Post{})

}
