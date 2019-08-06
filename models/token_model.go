package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TokenModel struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
	Time   time.Timer
}
