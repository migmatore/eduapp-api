package models

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	//Posts 	  []PostModel `json:"posts"`
}

type UserView struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"nick_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	//Posts 	  []PostModel `json:"posts"`
}
