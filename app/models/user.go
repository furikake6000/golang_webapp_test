package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	TwitterId  string
	ScreenName string
	ImageURL   string
}
