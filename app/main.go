package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"./controllers"
	"./secret"
)

func main() {
	router := gin.Default()
	router.Static("/assets/css", "./assets/css")
	router.Static("/assets/javascripts", "./assets/javascripts")
	router.LoadHTMLGlob("templates/*.html")

	// Settings for cookie
	store := sessions.NewCookieStore([]byte("tmp_secret_key"))
	router.Use(sessions.Sessions("GolangWebappTest", store))

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"tw_key":    secret.Credentials["twitter_key"],
			"tw_secret": secret.Credentials["twitter_secret"],
		})
	})

	// Login with Twitter
	router.GET("/auth/twitter", controllers.LoginByTwitter)
	router.GET("/auth/twitter/callback", controllers.TwitterCallback)

	router.Run(":8080")
}
