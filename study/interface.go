package main 
/*
Go 语言提供了另外一种数据类型即接口，
它把所有的具有共性的方法定义在一起，
任何其他类型只要实现了这些方法就是实现了这个接口。

定义接口
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

定义结构体
type struct_name struct {
    variables
}

实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
   方法实现 
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
   方法实现
}
*/
import (
	"fmt"
)

func main() {
	var phone Phone;
	phone = new(NokiaPhone);
	phone.call();

	phone = new(IPhone);
	phone.call();	
}

type Phone interface {
	call();//打电话
	sms();//发短信
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!");
}

func (nokiaPhone NokiaPhone) sms() {
	fmt.Println("I am Nokia, I can sms you!");
}

type IPhone struct {
}

func (iphone IPhone) call() {
	fmt.Println("I am IPhone, I can call you!");
}

func (iphone IPhone) sms() {
	fmt.Println("I am IPhone, I can sms you!");
}
