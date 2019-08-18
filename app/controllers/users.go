package controllers

import (
	"my/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	user := models.CurrentUser(c)

	c.HTML(200, "index.html", gin.H{
		"user": user,
	})
}
