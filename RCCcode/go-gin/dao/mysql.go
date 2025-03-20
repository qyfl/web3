package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitSQL() (err error) {
	dst := "root@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dst)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
