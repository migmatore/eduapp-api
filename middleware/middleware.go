package middleware

import (
	"fmt"
	"strings"
	"time"
	"user/eduAppApi/db"
	"user/eduAppApi/models"

	//jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SingingKey = "secret"

func GenerateToken(userId uint, login string, accessLevel string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":        login,
		"access_level": accessLevel,
		"exp":          time.Now().Add(time.Minute * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SingingKey))

	_token := models.TokenModel{
		UserId:      userId,
		Token:       tokenString,
		AccessLevel: accessLevel,
	}

	db.DB.Create(&_token)

	return tokenString, err
}

func MiddlewareHandler(accessLevel ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(SingingKey), nil
		})
		if err != nil {
			fmt.Errorf("Token not found", err.Error())

			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Token not found",
				"error":   err.Error(),
			})

			c.Abort()

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var _accessLevels []string

			for i, j := range claims {
				if i == "access_level" {
					_splitAccessLevels := strings.Split(j.(string), ", ")
					fmt.Println(_splitAccessLevels)

					for _, i := range _splitAccessLevels {
						_accessLevels = append(_accessLevels, i)
					}

					fmt.Println(_accessLevels)
				}
			}

			if len(accessLevel) > 1 {
				for i := range _accessLevels {
					if _accessLevels[i] == accessLevel[i] {
						c.JSON(http.StatusOK, gin.H{
							"status":  http.StatusOK,
							"message": "Access is allowed",
						})

						c.Next()

						return
					}
				}

				c.JSON(http.StatusLocked, gin.H{
					"status":  http.StatusLocked,
					"message": "You must have access levels: " + accessLevel[0] + ", " + accessLevel[1],
				})
			} else {
				for i := range _accessLevels {
					if _accessLevels[i] == accessLevel[0] {
						c.JSON(http.StatusOK, gin.H{
							"status":  http.StatusOK,
							"message": "Access is allowed",
						})

						c.Next()

						return
					}
				}

				c.JSON(http.StatusLocked, gin.H{
					"status":  http.StatusLocked,
					"message": "You must have access levels: " + accessLevel[0],
				})
			}

			c.Abort()

			return
		} else {
			fmt.Println(err)

			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "user needs to be signed",
			})

			c.Abort()

			return
		}
	}
}
