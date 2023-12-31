package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mo7rex/api-gin/controllers"
	"github.com/mo7rex/api-gin/database"
	"github.com/mo7rex/api-gin/middleware"
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
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.LogOut)
	r.DELETE("/account/:id", controllers.DeleteAccount)

	r.GET("/auth", middleware.RequireAuth, controllers.Validate)
	r.Run()

}
