package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/youhao?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return db
}
