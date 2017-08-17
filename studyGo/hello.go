//表示一个可独立执行的程序
package main 
//告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import "fmt"

func main() {
	//三种声明变量的方式
	//出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误
	var a int = 10;
	var b = 10;
	c := 10;
	/* 这是我的一个简单的程序 */
	fmt.Println("Hello GO",a,b,c);
}