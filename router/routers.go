package router

import (
	"github.com/gin-gonic/gin"
	"user/eduAppApi/handlers"
	"user/eduAppApi/middleware"
)

func SetupRouter() *gin.Engine {
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
			user.PUT("/update", handlers.UpdateUserHandler)
			user.DELETE("/delete", handlers.DeleteUserHandler)
		}

		post := api.Group("/post")
		post.Use(middleware.MiddlewareHandler("admin", "user"))
		{

			post.GET("/")
			post.POST("/create")
		}
	}

	return r
}
