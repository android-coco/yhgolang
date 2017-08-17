package main

import (
	"fmt"
	"awesomeProject/org.yh/util"//引用自己的包
)

type NokiaPhone struct {
}

type Iphone struct {
}

type I_Phone interface {
	call();
	sms();
}

func (iphone Iphone) call() {
	fmt.Println("my is Iphone call");
}

func (iphone NokiaPhone) sms() {
	fmt.Println("my is NokiaPhone sms");
}

func (iphone NokiaPhone) call() {
	fmt.Println("my is NokiaPhone call");
}

func (iphone Iphone) sms() {
	fmt.Println("my is Iphone sms");
}

func main() {
	var iphone I_Phone;
	iphone = new(NokiaPhone);
	iphone.call();
	iphone.sms();

	iphone = new(Iphone);
	iphone.call();
	iphone.sms();

	a := 100;
	b := 200;

	fmt.Print("转换前",a,b);
	fmt.Print("\n");
	util.Swap(&a,&b);
	fmt.Print("转换后",a,b);
	fmt.Print("\n");

	//定义一个切片
	//1,类型  2,长度  3,预计长度
	c := make([]int,3,8);//数组不能用make
	d := [3]int{};//数组
	e := []int{};//没有限定长度的数组就是切片
	util.PrintSlice(c);
	//util.PrintSlice(d);//报错  类型不一致
	for i, v := range d {
		fmt.Printf("[%d->%d]\n", i,v);
	}
	util.PrintSlice(e);
	//map
	f := map[string]string{"Youhao":"Lyl"};//一步到位初始化到赋值

	var f1 map[string]string;//定义
	f1 = make(map[string]string);//初始化
	f1["Youhao"] = "Lyl";//赋值

	//打印map
	for v := range f {
		fmt.Printf("[%s => %s]\n", v,f[v]);
	}
	for v := range f1 {
		fmt.Printf("[%s => %s]\n", v,f1[v]);
	}
}
