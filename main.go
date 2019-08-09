package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"user/eduAppApi/db"
	"user/eduAppApi/handlers"
	"user/eduAppApi/middleware"
)

func main() {
	db.InitMigration()

	//middleware, err := jwt.New(&jwt.GinJWTMiddleware{
	//	Key:         []byte(SIGNING_KEY),
	//	SigningAlgorithm: "HS256",
	//	Authenticator: func(c *gin.Context) (interface{}, error) {
	//		return []byte(SIGNING_KEY), nil
	//	},
	//})
	//if err != nil {
	//	panic(err.Error())
	//}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		api.GET("/admin")

		user := api.Group("/user")
		{
			user.POST("/create", handlers.CreateUserHandler)
			user.GET("/login", handlers.LoginUserHandler)
			user.PUT("/update", handlers.UpdateUserHandler).Use(middleware.MiddlewareHandler("user"))
			user.DELETE("/delete", handlers.DeleteUserHandler).Use(middleware.MiddlewareHandler("user"))
		}

		post := api.Group("/post")
		post.Use(middleware.MiddlewareHandler("admin"))
		{
			post.GET("/", getPostsHandler)
			post.POST("/")
		}
	}

	_ = r.Run()
}

func getPostsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "token is got",
	})
}