package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mo7rex/api-gin/database"
	"github.com/mo7rex/api-gin/model"
)

func RequireAuth(ctx *gin.Context) {
	//check the cookie
	tokenString, err := ctx.Cookie("auth")
	//decode and validate
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		//find the user with the token sub
		var account model.Account
		database.DB.Find(&account, claims["sub"])
		if account.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach to req
		ctx.Set("account", account)
		//countinue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

}
