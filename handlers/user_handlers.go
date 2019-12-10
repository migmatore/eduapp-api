package handlers

import (
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
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	login := c.PostForm("login")
	password := c.PostForm("password")

	var users []models.UserModel
	var sortedUsers []string

	_user := models.UserModel{
		FirstName: firstName,
		LastName:  lastName,
		Login:     login,
		Password:  password,
		Role:      "user",
	}

	db.DB.Find(&users)

	for _, user := range users {
		sortedUsers = append(sortedUsers, strings.ToLower(user.Login))
	}

	if len(sortedUsers) <= 1 {
		if len(sortedUsers) == 0 {
			db.DB.Create(&_user)

			c.JSON(http.StatusCreated, gin.H{
				"status":  http.StatusCreated,
				"message": "user created",
			})

			return
		} else {
			for _, userLogin := range sortedUsers {
				if userLogin == strings.ToLower(login) {
					c.JSON(http.StatusNotFound, gin.H{
						"status":  http.StatusNotFound,
						"message": "user couldn't be created",
					})

					return
				} else {
					db.DB.Create(&_user)

					c.JSON(http.StatusCreated, gin.H{
						"status":  http.StatusCreated,
						"message": "user created",
					})

					return
				}
			}
		}
	} else {
		utils.SortStrings(sortedUsers)

		if sort.StringsAreSorted(sortedUsers) && !(0 > len(sortedUsers) || 0 == len(sortedUsers)) {
			if utils.BinSearchString(sortedUsers, strings.ToLower(login), 0, len(sortedUsers)) {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": "user couldn't be created",
				})

				return
			} else {
				db.DB.Create(&_user)

				c.JSON(http.StatusCreated, gin.H{
					"status":  http.StatusCreated,
					"message": "user created",
				})

				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "array not sorted",
		})

		return
	}
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
			"status":  http.StatusNotFound,
			"message": "Users not found!",
		})

		return
	}

	for _, user := range users {
		sortedUsers = append(sortedUsers, user.Login)
	}

	utils.SortStrings(sortedUsers)

	if sort.StringsAreSorted(sortedUsers) && !(0 > len(sortedUsers)) {
		search := utils.BinSearchString(sortedUsers, login, 0, len(sortedUsers))

		if search {
			db.DB.Where("login = ?", login).First(&user)

			if password == user.Password {
				token, err := middleware.GenerateToken(user.ID, login, user.Role)
				if err != nil {
					panic(err.Error())
				}

				c.JSON(http.StatusCreated, gin.H{
					"status":  http.StatusCreated,
					"message": "token created",
					"token":   token,
				})

				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "user not found",
		})

		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": "user not found",
	})

	return
}

func UpdateUserHandler(c *gin.Context) {
	id := c.PostForm("id")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	login := c.PostForm("login")
	password := c.PostForm("password")

	var user models.UserModel

	db.DB.First(&user, id)

	db.DB.Model(&user).Updates(models.UserModel{
		FirstName: firstName,
		LastName:  lastName,
		Login:     login,
		Password:  password,
		Role:      "user",
	})

	db.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user was updated",
	})

	return
}

func DeleteUserHandler(c *gin.Context) {
	userId := c.Query("id")

	var user models.UserModel

	db.DB.Unscoped().Delete(&user, userId)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user deleted",
	})

	return
}
