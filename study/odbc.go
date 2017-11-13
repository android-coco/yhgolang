package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lunny/godbc"
)

type NxServerState struct {
	ID             int       `xorm:"pk not null 'ID'"`
	GameID         int       `xorm:"not null 'GameID'"`
	IssuerId       int       `xorm:"not null IssuerId"`
	ServerID       int       `xorm:"not null ServerID"`
	ServerName     string    `xorm:"ServerName"`
	OnlineNum      int       `xorm:"not null OnlineNum"`
	MaxOnlineNum   int       `xorm:"not null MaxOnlineNum"`
	ServerIP       string    `xorm:"not null ServerIP"`
	Port           int       `xorm:"not null Port"`
	IsRuning       int       `xorm:"not null IsRuning"`
	ServerStyle    int       `xorm:"ServerStyle"`
	IsStartIPWhile int       `xorm:"not null IsStartIPWhile"`
	LogTime        time.Time `xorm:"IsStartIPWhile"`
	UpdateTime     time.Time `xorm:"UpdateTime"`
	OrderBy        int       `xorm:"not null OrderBy"`
}

func main() {
	File, _ := os.Create("result")
	defer File.Close()
	Engine, err := xorm.NewEngine("odbc", "driver={SQL Server};Server=127.0.0.1;Database=fuck;uid=sa;pwd=123456;")
	if err != nil {
		fmt.Println("新建引擎", err)
		return
	}
	if err := Engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	Engine.SetTableMapper(core.SameMapper{})
	Engine.ShowSQL = true
	Engine.SetMaxConns(5)
	Engine.SetMaxIdleConns(5)
	result := new(NxServerState)
	lines, _ := Engine.Rows(result)
	defer lines.Close()
	lines.Next()
	r := new(NxServerState)
	for {
		err = lines.Scan(r)
		if err != nil {
			return
		}
		fmt.Println(*r)
		File.WriteString(fmt.Sprintln(*r))
		if !lines.Next() {
			break
		}
	}
}
