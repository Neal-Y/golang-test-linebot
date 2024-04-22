package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

func InitMySQL() (err error) {
	for i := 0; i < 10; i++ {
		Db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: "admin:1234@tcp(localhost:3306)/be103?charset=utf8&parseTime=True&loc=Local",
		}), &gorm.Config{})

		if err == nil {
			return
		}
		time.Sleep(2 * time.Second)
	}
	return
}
