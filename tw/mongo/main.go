package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
)

var db *gorm.DB

func init() {
	//rm-wz9ba43n1i6sfvt1p35930.mysql.rds.aliyuncs.com:3306/fm
	var err error
	db, err = gorm.Open("mysql", "pc:uJREJW9DNIk2H3I96ayz@tcp(rm-wz9ba43n1i6sfvt1p35930.mysql.rds.aliyuncs.com:3306)/fm?charset=utf8&parseTime=true&loc=Local")
	//defer db.Close()
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return db
}

type Task struct {
	Token string `gorm:"token" json:"token"`
}

func (Task) TableName() string {
	return "biz_task"
}

//120.78.69.55:27018
func main() {
	var list []Task
	err := GetDb().Model(&Task{}).
		Where("code = ? AND  create_time >= ? AND create_time <= ?", "SUCCESS", "2021-06-01 00:00:00", "2021-06-30 23:59:59").
		Order("create_time desc").Find(&list).Error
	if len(list) == 0 {
		fmt.Println("查询token失败")
		return
	}
	// TW 老db
	session, err := mgo.Dial("172.18.215.127:27018")
	//session, err := mgo.Dial("120.78.69.55:27018")
	if err != nil {
		fmt.Println("=====1",err)
		panic(err)
	}
	defer session.Close()

	//新DB
	session1, err1 := mgo.Dial("47.106.74.103:27018")
	if err1 != nil {
		fmt.Println("=====2",err1)
		panic(err1)
	}
	defer session1.Close()

	session.SetMode(mgo.Monotonic, true)
	session1.SetMode(mgo.Monotonic, true)

	for key, v := range list {
		if v.Token != "" {
			paramData := bson.M{}
			paramData["token"] = v.Token
			resultData := bson.M{}
			err1 := session.DB("mh-data").C("mhData").Find(paramData).One(resultData)
			if err1 != nil {
				fmt.Println("err1:", key, v.Token, err1)
				continue
			}
			err2 := session1.DB("mh-data").C("mhData").Insert(resultData)
			if err2 != nil {
				fmt.Println("err2:", key, v.Token, err2)
				continue
			}

			paramLog := bson.M{}
			paramLog["token"] = v.Token
			resultLog := bson.M{}
			err3 := session.DB("mh-data").C("mhLog").Find(paramLog).One(resultLog)
			if err3 != nil {
				fmt.Println("err3:", key, v.Token, err3)
				continue
			}
			err4 := session1.DB("mh-data").C("mhLog").Insert(resultLog)
			if err4 != nil {
				fmt.Println("err4:", key, v.Token, err4)
				continue
			}
			fmt.Println(key, v.Token)
		}
	}
}
