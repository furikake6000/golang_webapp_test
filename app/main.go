package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"./controllers"
	"./models"
	"./secret"
)

func main() {
	router := gin.Default()
	var db *gorm.DB
	for {
		var err error
		db, err = gorm.Open("mysql", "root:mysql@tcp(db)/golang_webapp")

		// 接続成功するまでポーリング
		if err == nil {
			break
		}
	}
	router.LoadHTMLGlob("templates/*.html")

	// Settings for cookie
	store := sessions.NewCookieStore([]byte("tmp_secret_key"))
	router.Use(sessions.Sessions("GolangWebappTest", store))

	db.AutoMigrate(&models.User{})

	router.GET("/", func(cont *gin.Context) {
		cont.HTML(200, "index.html", gin.H{
			"tw_key":    secret.Credentials["twitter_key"],
			"tw_secret": secret.Credentials["twitter_secret"],
		})
	})

	// Login with Twitter
	router.GET("/auth/twitter", controllers.LoginByTwitter)
	router.GET("/auth/twitter/callback", controllers.TwitterCallback)

	router.Run(":8080")

	defer db.Close()
}
