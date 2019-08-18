package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"my/controllers"
)

func main() {
	router := gin.Default()
	router.Static("/assets/css", "./assets/css")
	router.Static("/assets/javascripts", "./assets/javascripts")
	router.LoadHTMLGlob("templates/*.html")

	// Settings for cookie
	store := sessions.NewCookieStore([]byte("tmp_secret_key"))
	router.Use(sessions.Sessions("GolangWebappTest", store))

	// Root
	router.GET("/", controllers.Index)

	// Trends
	router.GET("/trend", controllers.ShowTrend)

	// Login with Twitter
	router.GET("/auth/twitter", controllers.LoginByTwitter)
	router.GET("/auth/twitter/callback", controllers.TwitterCallback)

	router.Run(":8080")
}
