package models

import "github.com/jinzhu/gorm"

type PostModel struct {
	gorm.Model
	Title string     `json:"title"`
	Body  string     `json:"body"`
	Tags  []TagModel `json:"tags" gorm:"foreignkey:ID"`
}
