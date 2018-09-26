package main

import (
	"yhgolang/gorm/db"
	"fmt"
)

type Loss struct {
	StoreID  int64
	Dateline string
	Type     int64
}

func (*Loss) TableName() string {
	 return "loss"
}

func main() {
	db := db.GetDb()
	db.LogMode(true)
	loss := Loss{}
	if err := db.First(&loss).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n",loss)
}
