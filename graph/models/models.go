package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int    `json:"userId"`
	User   *User  `json:"user"`
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}
