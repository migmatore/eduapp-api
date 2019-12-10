package models

import "github.com/jinzhu/gorm"

type TagModel struct {
	gorm.Model
	Title string `json:"title"`
}

type TagView struct {
	Id    uint   `json:"id"`
	Title string `json:"title"`
}
