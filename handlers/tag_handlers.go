package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
	"user/eduAppApi/db"
	"user/eduAppApi/models"
	"user/eduAppApi/utils"
)

func GetTagsHandler(c *gin.Context) {
	var tags []models.TagModel
	var _tags []models.TagView

	db.DB.Find(&tags)

	if len(tags) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Tags not found",
		})

		return
	}

	for _, tag := range tags {
		_tags = append(_tags, models.TagView{
			Id:    tag.ID,
			Title: tag.Title,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": _tags,
	})

	return
}

func CreateTagHandler(c *gin.Context) {
	title := c.PostForm("title")

	var tags []models.TagModel
	var sortedTags []string

	_tag := models.TagModel{
		Title: title,
	}

	db.DB.Find(&tags)

	for _, tag := range tags {
		sortedTags = append(sortedTags, strings.ToLower(tag.Title))
	}

	if len(sortedTags) <= 1 {
		if len(sortedTags) == 0 {
			db.DB.Create(&_tag)

			c.JSON(http.StatusCreated, gin.H{
				"status":  http.StatusOK,
				"message": "Tag created",
			})

			return
		} else {
			for _, tagTitle := range sortedTags {
				if tagTitle == strings.ToLower(title) {
					c.JSON(http.StatusNotFound, gin.H{
						"status":  http.StatusNotFound,
						"message": "tag couldn't be created",
					})

					return
				} else {
					db.DB.Create(&_tag)

					c.JSON(http.StatusCreated, gin.H{
						"status":  http.StatusCreated,
						"message": "Tag created",
					})

					return
				}
			}
		}
	} else {
		utils.SortStrings(sortedTags)

		if sort.StringsAreSorted(sortedTags) && !(0 > len(sortedTags) || 0 == len(sortedTags)) {
			if utils.BinSearchString(sortedTags, strings.ToLower(title), 0, len(sortedTags)) {
				c.JSON(http.StatusNotFound, gin.H{
					"status":  http.StatusNotFound,
					"message": "tag couldn't be created",
				})

				return
			} else {
				db.DB.Create(&_tag)

				c.JSON(http.StatusCreated, gin.H{
					"status":  http.StatusCreated,
					"message": "tag created",
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

func EditTagHandler(c *gin.Context) {

}

func DeleteTagHandler(c *gin.Context) {

}
