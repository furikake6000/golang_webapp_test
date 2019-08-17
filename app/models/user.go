package models

import (
	"../db"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `gorm:"unique;not null"`
	TwitterId  string
	ScreenName string
	ImageURL   string
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&User{})
}
