package models

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	NickName  string      `json:"nick_name"`
	Password  string      `json:"password"`
	//Posts 	  []PostModel `json:"posts"`
}

type UserView struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	NickName  string      `json:"nick_name"`
	Password  string      `json:"password"`
	//Posts 	  []PostModel `json:"posts"`
}
