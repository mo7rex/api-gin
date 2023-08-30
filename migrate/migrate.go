package main

import (
	"fmt"
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
	err := database.DB.AutoMigrate(&model.Account{}, &model.Post{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// fmt.Println(database.DB.Migrator().CreateConstraint(&model.Account{}, "Post"))
	// fmt.Println(database.DB.Migrator().CreateConstraint(&model.Account{}, "fk_accounts_post"))
	// fmt.Println(database.DB.Migrator().HasConstraint(&model.Account{}, "Post"))
	// fmt.Println(database.DB.Migrator().HasConstraint(&model.Account{}, "fk_accounts_post"))

}
