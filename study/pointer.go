package main 

import (
   "fmt"
)
const MAX int = 3;//常量
func main() {
	var a int = 20;   /* 声明实际变量 */
    var ip *int;        /* 声明指针变量 */

    ip = &a;  /* 指针变量的存储地址 */

    fmt.Printf("a 变量的地址是: %x1\n", &a  );

    /* 指针变量的存储地址 */
    fmt.Printf("ip 变量储存的指针地址: %x1\n", ip );

    /* 使用指针访问值 */
    fmt.Printf("*ip 变量的值: %d\n", *ip );

    fmt.Printf("============================\n");

    var ipx *int;

    fmt.Printf("ipx 的值为 : %x1\n", ipx);
    fmt.Printf("ipx 的值为 : %t\n", ipx == nil);

    fmt.Printf("============指针数组===============\n");
    jj := []int{10,100,200};
    //jj := [...]int{10,100,200};
    var i int;

    for i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i, jj[i] )
    }


    a1 := []int{10,100,200}
   	var i1 int;
   	var ptr [MAX]*int;

   	for  i1 = 0; i1 < MAX; i1++ {
      	ptr[i1] = &a1[i1]; /* 整数地址赋值给指针数组 */
   	}

   	for  i1 = 0; i1 < MAX; i1++ {
       fmt.Printf("a[%d] = %d\n", i1, *ptr[i1]);
   	}

    fmt.Printf("============指针的指针===============\n");

    var c int;
    var ptr1 *int;
    var pptr1 **int;

    c = 3000;

	/* 指针 ptr 地址 */
    ptr1 = &c;

	/* 指向指针 ptr 地址 */
    pptr1 = &ptr1;

   fmt.Printf("变量 c = %d\n", c );
   fmt.Printf("指针变量 *ptr1 = %d\n", *ptr1 );
   fmt.Printf("指向指针的指针变量 **pptr1 = %d\n", **pptr1);

   fmt.Printf("============指针作为函数参数===============\n");

   /* 定义局部变量 */
   var x int = 100;
   var y int = 200;

   fmt.Printf("替换前 x1 = %d\n", x);
   fmt.Printf("替换前 y1 = %d\n", y);

   /* 调用函数用于交换值
   * &x1 指向 x1 变量的地址
   * &y1 指向 y1 变量的地址
   */
   swap(&x , &y);

   fmt.Printf("替换后 x1 = %d\n", x);
   fmt.Printf("替换后 y1 = %d\n", y);

}

func swap(x *int , y *int) {
	var temp int;
	temp = *x;
	*x = *y;
	*y = temp;
}