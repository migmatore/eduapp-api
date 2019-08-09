package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
	"user/eduAppApi/db"
	"user/eduAppApi/middleware"
	"user/eduAppApi/models"
	"user/eduAppApi/utils"
)

func CreateUserHandler(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	var users []models.UserModel
	var sortedUsers []string

	_user := models.UserModel{
		Login: login,
		Password: password,
	}

	db.DB.Find(&users)

	for _, user := range users {
		sortedUsers = append(sortedUsers, strings.ToLower(user.Login))
	}

	utils.SortStrings(sortedUsers)

	if sort.StringsAreSorted(sortedUsers) && !(0 > len(sortedUsers)){
		if utils.BinSearchString(sortedUsers, strings.ToLower(login), 0, len(sortedUsers)) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"message": "user couldn't be created",
			})

			return
		} else {
			db.DB.Create(&_user)

			c.JSON(http.StatusCreated, gin.H{
				"status": http.StatusCreated,
				"message": "user created",
			})

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"message": "array not sorted",
	})
}

func LoginUserHandler(c *gin.Context) {
	login := c.Query("login")
	password := c.Query("password")

	var users []models.UserModel
	var user models.UserModel
	var sortedUsers []string

	db.DB.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Users not found!",
		})

		return
	}

	for _, user := range users {
		sortedUsers = append(sortedUsers, user.Login)
	}

	utils.SortStrings(sortedUsers)

	if sort.StringsAreSorted(sortedUsers) && !(0 > len(sortedUsers)){
		search := utils.BinSearchString(sortedUsers, login, 0, len(sortedUsers))

		if search {
			db.DB.Where("login = ?", login).First(&user)

			if password == user.Password {
				token, err := middleware.GenerateToken(user.ID, login, user.Role)
				if err != nil {
					panic(err.Error())
				}

				c.JSON(http.StatusCreated, gin.H{
					"status": http.StatusCreated,
					"message": "token created",
					"token": token,
				})

				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "array not sorted",
		})

		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"message": "user not found",
	})
}

// ****************
// TODO: UpdateUserHandler
// ****************
func UpdateUserHandler(c *gin.Context) {

}

func DeleteUserHandler(c *gin.Context) {
	userId := c.Query("id")

	var user models.UserModel

	err := db.DB.Unscoped().Delete(&user, userId)
	if err !=  nil {
		fmt.Errorf("Error: %v", err.Error)

		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "user not found",
			"error": err.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "user deleted",
	})
}
