package models

import (
	"../db"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	TwitterId  string `gorm:"unique;not null"`
	ScreenName string
	ImageURL   string
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&User{})
}
