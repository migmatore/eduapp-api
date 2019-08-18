package models

import (
	"github.com/jinzhu/gorm"
)

type TokenModel struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	Token       string `json:"token"`
	AccessLevel string `json:"access_level"`
}
