package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"user/eduAppApi/db"
	"user/eduAppApi/router"
)

func main() {
	db.InitMigration()

	r := router.SetupRouter()

	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}

func getPostsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "token is got",
	})
}
