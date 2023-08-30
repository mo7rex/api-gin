package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mo7rex/api-gin/controllers"
	"github.com/mo7rex/api-gin/database"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	database.ConnectToDatabase()

}

func main() {
	fmt.Println("Hello world")
	r := gin.Default()
	r.GET("/post", controllers.ShowAllPosts)
	r.GET("/post/:id", controllers.ShowOnePost)
	r.POST("/post", controllers.CreatePost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeleteOnePost)

	r.GET("/account/:account_id", controllers.GetAllAccountsWithPosts)
	r.POST("/account", controllers.CreateAccount)
	r.DELETE("/account/:id", controllers.DeleteAccount)
	r.Run()

}
