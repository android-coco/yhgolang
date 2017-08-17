//表示一个可独立执行的程序
package main 
//告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import "fmt"
import "unsafe"
var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int;
    b bool;
);
const YH = "yh.db";
var c, d int = 1, 2;
var e, f = 123, "hello";
//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
const (
    a1 = "abc";
    b2 = len(a1);
    c3 = unsafe.Sizeof(a1);
);
const (
    i = 1<<iota;
    j = 3<<iota;
    k;
    l;
);
func main() {
	//这是使用变量的首选形式，但是它只能被用在函数体内，
	//而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
	 g, h := 123, "hello";
	 var a1 string = "abc"
	/* 这是我的一个简单的程序 */
	 println(x, y, a, b, c, d, e, f, g, h,a1,YH);
	 //fmt.Println(x, y, a, b, c, d, e, f, g, h);
	
	   const LENGTH int = 10;
	   const WIDTH int = 5;
	   var area int;
	   const a, b, c = 1, false, "str"; //多重赋值
	   area = LENGTH * WIDTH;
	   fmt.Printf("面积为 : %d", area);
	   println();
	   println(a, b, c); 
	   println(a1, b2, c3);
	   fmt.Println("i=",i);
	   fmt.Println("j=",j);
	   fmt.Println("k=",k);
	   fmt.Println("l=",l);
}