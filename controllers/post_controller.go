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
	if body.Title == "" || body.Body == "" {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	post := model.Post{Title: body.Title, Body: body.Body}
	res := database.DB.Create(&post)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"post": post})
}

func ShowAllPosts(ctx *gin.Context) {
	var posts []model.Post
	res := database.DB.Find(&posts)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"post": posts})
}

func ShowOnePost(ctx *gin.Context) {
	var post model.Post
	id, _ := ctx.Params.Get("id")

	res := database.DB.First(&post, id)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"post": post})
}

func UpdatePost(ctx *gin.Context) {
	var input struct {
		Title string
		Body  string
	}
	ctx.Bind(&input)
	if input.Title == "" || input.Body == "" {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	var post model.Post
	id, _ := ctx.Params.Get("id")
	database.DB.First(&post, id)
	res := database.DB.Model(&post).Updates(model.Post{Title: input.Title, Body: input.Body})
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"post": post})
}

func DeleteOnePost(ctx *gin.Context) {
	var post model.Post
	id, err := ctx.Params.Get("id")
	if !err {
		ctx.Status(404)
	}
	res := database.DB.First(&post, id).Delete(&post)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"OK": "The post is deleted"})
}
