package main

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  gorm.Model
  Name string
  Age int
}

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

  router.GET("/", func(cont *gin.Context) {
    cont.HTML(200, "index.html", gin.H{})
  })
  
  db.AutoMigrate(&User{})

  router.Run(":8080")

  defer db.Close()
}
