package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mo7rex/api-gin/database"
	"github.com/mo7rex/api-gin/model"
)

func CreateAccount(ctx *gin.Context) {
	var input struct {
		Full_name    string
		Phone_number string
		Gender       string
	}
	ctx.Bind(&input)
	if input.Full_name == "" || input.Gender == "" || input.Phone_number == "" {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	account := model.Account{FullName: input.Full_name, PhoneNumber: input.Phone_number, Gender: input.Gender}
	res := database.DB.Create(&account)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"Account": account})

}

func GetAllAccount(ctx *gin.Context) {

	var accounts []model.Account
	res := database.DB.Find(&accounts)
	if res.Error != nil {
		ctx.Status(404)
		ctx.JSON(404, gin.H{"Error": "Not found!"})
		return
	}
	ctx.JSON(200, gin.H{"Accounts": accounts})
}
func GetAllAccountsWithPosts(ctx *gin.Context) {
	//create post, account instante
	var post []*model.Post
	var account model.Account
	//get the account id from the path
	acc_id, _ := ctx.Params.Get("account_id")
	//get the data from database
	res := database.DB.Find(&account, acc_id)
	database.DB.Where("account_id= ?", acc_id).First(&post)
	//append the post
	account.Post = post
	//check the error
	if res.Error != nil {
		ctx.Status(404)
		ctx.JSON(404, gin.H{"Error": "Not found!"})
		return
	}
	//decode
	ctx.JSON(200, gin.H{"Accounts": account})
}

func DeleteAccount(ctx *gin.Context) {
	var post []model.Post
	var account model.Account
	acc_id, found := ctx.Params.Get("id")
	if !found {
		ctx.Status(404)
		ctx.JSON(404, gin.H{"ERROR": "Not found"})
		return
	}
	database.DB.Select("account_id").Where("account_id = ?", acc_id).Delete(&post)
	res := database.DB.First(&account, acc_id).Delete(&account)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"OK": "The account is deleted"})
}
