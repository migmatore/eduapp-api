package main

import (
	"time"
	"user/eduAppApi/models"

	//jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SIGNING_KEY = "secret"

func loginUserHandler(c *gin.Context) {
	nickName := c.Query("nick_name")
	password := c.Query("password")

	var users []models.UserModel

	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Users not found!",
		})

		return
	}

	for _, user := range users {
		if (user.NickName == nickName) && (user.Password == password) {
			token, err := generateToken(nickName)
			if err != nil {
				panic(err.Error())
			}

			_token := models.TokenModel{
				UserId: user.ID,
				Token: token,
			}

			db.Create(&_token)

			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"message": "token created",
				"token": token,
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"message": "user not found",
	})
}

func generateToken(nickName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nick_name": nickName,
		"exp": time.Now().Add(time.Minute * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SIGNING_KEY))

	return tokenString, err
}
