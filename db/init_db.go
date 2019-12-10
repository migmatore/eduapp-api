package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"user/eduAppApi/models"
)

var DB *gorm.DB

var dbUrl = ""
var dbUrlDev = "host=localhost port=5432 user=admin dbname=testEduAppApi password=admin sslmode=disable"

func InitMigration() {
	var err error
	DB, err = gorm.Open("postgres", dbUrlDev) //sslmode=disable

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.UserModel{})
	DB.AutoMigrate(&models.TokenModel{})
	DB.AutoMigrate(&models.PostModel{})
	DB.AutoMigrate(&models.TagModel{})
}
