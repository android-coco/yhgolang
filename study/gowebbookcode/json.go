package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

var str = `
{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}
`

type Server struct {
	ServerName string `json:"servername" xorm:"extends"`
	ServerIP   string `json:"serverip"`
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	s1 := Server{ServerName: "1", ServerIP: "192.168.0.1"}
	j, _ := json.Marshal(s1)
	fmt.Println(string(j))

	// 获取tag中的内容
	t := reflect.TypeOf(&s1)
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json"))
	// 输出：servername
	fmt.Println(field.Tag.Get("xorm"))
	// 输出：extends
	fmt.Println(field.Interface())
}
