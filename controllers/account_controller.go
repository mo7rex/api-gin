package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mo7rex/api-gin/database"
	"github.com/mo7rex/api-gin/model"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var input struct {
		Email        string
		Password     string
		Full_name    string
		Phone_number string
		Gender       string
	}
	err := ctx.Bind(&input)
	if err != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	//hash the pass
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
	}
	account := model.Account{FullName: input.Full_name, PhoneNumber: input.Phone_number, Gender: input.Gender, Email: input.Email, Password: string(hash)}
	res := database.DB.Create(&account)
	if res.Error != nil {
		ctx.Status(400)
		ctx.JSON(400, gin.H{"Error": "Bad Request"})
		return
	}
	ctx.JSON(200, gin.H{"Account": account})

}

func Login(ctx *gin.Context) {
	var input struct {
		Email    string
		Password string
	}
	err := ctx.Bind(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	//check the email
	var account model.Account
	database.DB.Find(&account, "email= ?", input.Email)
	if account.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error!": "Invalid email"})
		return
	}
	//compare the sent password with the hash
	error := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(input.Password))
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Password"})
		return
	}
	//genrate the jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "error in genrate the token"})
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("auth", tokenstring, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"200": "OK",
	})

}
func LogOut(ctx *gin.Context) {
	ctx.SetCookie("auth", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"200": "OK",
	})
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

func Validate(ctx *gin.Context) {
	account, _ := ctx.Get("account")

	ctx.JSON(http.StatusOK, gin.H{"account": account})
}
