package models

import (
	"my/db"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
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

func CurrentUser(c *gin.Context) *User {
	session := sessions.Default(c)
	twitter_uid := session.Get("twitter_uid")
	if twitter_uid == nil {
		return nil
	}

	db := db.GetDB()

	user := &User{}
	db.Where(User{TwitterId: twitter_uid.(string)}).First(&user)
	return user
}
