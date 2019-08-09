package middleware

import (
	"fmt"
	"time"
	"user/eduAppApi/db"
	"user/eduAppApi/models"

	//jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SIGNING_KEY = "secret"

func GenerateToken(userId uint, login string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"role": role,
		"exp": time.Now().Add(time.Minute * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SIGNING_KEY))

	_token := models.TokenModel{
		UserId: userId,
		Token: tokenString,
		Role: role,
	}

	db.DB.Create(&_token)

	return tokenString, err
}

func MiddlewareHandler(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(SIGNING_KEY), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == role[0] {
				c.Next()
			} else if claims["role"] == role[0] && claims["role"] == role[1] {
				c.Next()
			} else {
				fmt.Println("Users isn't admin")
				c.Abort()

				return
			}

		} else {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "user needs to be signed",
			})
			c.Abort()

			return
		}
	}
}
