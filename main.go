package main

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"user/eduAppApi/models"
)

var db *gorm.DB

var dbUrl = ""
var dbUrlDev = "host=localhost port=5432 user=admin dbname=testEduAppApi password=admin sslmode=disable"

func initMigration() {
	var err error
	db, err = gorm.Open("postgres", dbUrlDev) //sslmode=disable

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserModel{})
	db.AutoMigrate(&models.TokenModel{})
	db.AutoMigrate(&models.PostModel{})
}

func main() {
	initMigration()

	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(SIGNING_KEY),
		SigningAlgorithm: "HS256",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return []byte(SIGNING_KEY), nil
		},
	})
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/admin")

		user := api.Group("/user")
		{
			user.POST("/create", createUser)
			user.GET("/login", loginUserHandler)
		}

		post := api.Group("/post")
		post.Use(middleware.MiddlewareFunc())
		{
			post.GET("/", getPostsHandler)
			post.POST("/")
		}
	}

	_ = r.Run()
}

func createUser(c *gin.Context) {
	nickName := c.PostForm("nick_name")
	password := c.PostForm("password")

	_user := models.UserModel{
		NickName: nickName,
		Password: password,
	}

	db.Create(&_user)
}

func getPostsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "token is got",
	})
}