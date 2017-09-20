package main

import (
	"fmt"
)

type MsgType int

type MessageCreator func() AnyMessage

type AnyMessage interface {
	UnmarshalBinary(data []byte) (n int, err error)
	MarshalBinary() (data []byte, err error)
}

type spec struct {
	name    string
	creator func() AnyMessage
}

type LoginMessage struct {
	user string
	pass string
	err  error
}

func (loginMessge LoginMessage) UnmarshalBinary(data []byte) (n int, err error) {
	fmt.Println("UnmarshalBinary")
	return 1, nil
}

func (loginMessge LoginMessage) MarshalBinary() (data []byte, err error) {
	fmt.Println("MarshalBinary")
	return []byte(loginMessge.user), nil
}

//登录
func NewLoginMessage() AnyMessage {
	loginMessage := LoginMessage{"yh", "123456", fmt.Errorf("%s", "登录成功")}
	return loginMessage
}

//注销
func LoginOut() AnyMessage {
	loginMessage := LoginMessage{"yh", "123456", fmt.Errorf("%s", "退出成功")}
	return loginMessage
}

var (
	msgTypeSpec = map[MsgType]spec{
		1: spec{"登录消息", NewLoginMessage}, //这里存函数名称
		2: spec{"注销消息", LoginOut},        //这里存函数名称
	}
)

func (mt *MsgType) Name() string {
	return msgTypeSpec[*mt].name
}

func (mt *MsgType) CreateMessage() AnyMessage {
	entry, ok := msgTypeSpec[*mt]
	if !ok {
		return nil
	}

	if entry.creator == nil {
		return nil
	}
	return entry.creator() //形成闭包(函数调用)
}

func main() {
	var mt MsgType = 2 //根据类型判断登录状态
	msg := mt.CreateMessage()
	if login, ok := msg.(LoginMessage); ok {
		fmt.Printf("登录名称:%s  登录密码:%s  错误码:%s", login.user, login.pass, login.err.Error())
	}
	//fmt.Printf("msg:%v\n", msg)
}
