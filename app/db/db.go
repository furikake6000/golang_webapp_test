package db

import "github.com/jinzhu/gorm"

var db *gorm.DB

func init() {
	for {
		var err error
		db, err = gorm.Open("mysql", "root:mysql@tcp(db)/golang_webapp?parseTime=true")

		// 接続成功するまでポーリング
		if err == nil {
			break
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
