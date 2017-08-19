package main

import (
	"fmt"
)

type empty interface {
}

type USB interface {
	Name() string
	//Connect()
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	//var a USB
	a := PhoneConnecter{name: "PhoneConnecter"}
	a.Connect()
	Disconnect(a)
	Disconnect1(a)
	//注意事项  接口存的是拷贝  无法修改内容
	//只有当接口存储的类型和对象都为nil时,接口才是nil
	//空接口可以作为任何类型数据的容器
	var b interface{}
	fmt.Println(b == nil)

	var p *int = nil
	b = p
	fmt.Println(b == nil)
}

//interface{} 空接口  Object
func Disconnect(usb USB) {
	//判断是否不是那个同一个类型usb.(PhoneConnecter)
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected:", pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}

func Disconnect1(usb interface{}) {
	//type switch  fallthrough不能用
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}

}
