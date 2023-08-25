package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mo7rex/api-gin/database"
	"github.com/mo7rex/api-gin/model"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	ctx.Bind(&body)
	post := model.Post{Title: body.Title, Body: body.Body}
	res := database.DB.Create(&post)
	if res.Error != nil {
		ctx.Status(400)
	}
	ctx.JSON(200, gin.H{"post": post})
}
