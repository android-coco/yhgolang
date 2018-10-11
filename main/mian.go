package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"yhgolang/org.yh/util" //引用自己的包
)

type NokiaPhone struct {
}

type Iphone struct {
}

type I_Phone interface {
	call()
	sms()
}

func (iphone Iphone) call() {
	fmt.Println("my is Iphone call")
}

func (iphone NokiaPhone) sms() {
	fmt.Println("my is NokiaPhone sms")
}

func (iphone NokiaPhone) call() {
	fmt.Println("my is NokiaPhone call")
}

func (iphone Iphone) sms() {
	fmt.Println("my is Iphone sms")
}

//排序
type Person struct {
	Name string
	Age  int64
}
type Persons []Person

// 获取此 slice 的长度
func (p Persons) Len() int { return len(p) }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p Persons) Less(i, j int) bool {
	return p[i].Age > p[j].Age
}

// 交换数据
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	fmt.Printf("pid= %#v\n", os.Getpid())
	var iphone I_Phone
	iphone = new(NokiaPhone)
	iphone.call()
	iphone.sms()

	iphone = new(Iphone)
	iphone.call()
	iphone.sms()

	a := 100
	b := 200

	fmt.Print("转换前", a, b)
	fmt.Print("\n")
	util.Swap(&a, &b)
	fmt.Print("转换后", a, b)
	fmt.Print("\n")

	//定义一个切片
	//1,类型  2,长度  3,预计长度
	c := make([]int, 3, 8) //数组不能用make
	d := [3]int{}          //数组
	e := []int{}           //没有限定长度的数组就是切片
	util.PrintSlice(c)
	//util.PrintSlice(d);//报错  类型不一致
	for i, v := range d {
		fmt.Printf("[%d->%d]\n", i, v)
	}
	util.PrintSlice(e)
	//map
	f := map[string]string{"Youhao": "Lyl"} //一步到位初始化到赋值

	var f1 map[string]string     //定义
	f1 = make(map[string]string) //初始化
	f1["Youhao"] = "Lyl"         //赋值

	//打印map
	for v := range f {
		fmt.Printf("[%s => %s]\n", v, f[v])
	}
	for v := range f1 {
		fmt.Printf("[%s => %s]\n", v, f1[v])
	}

	fmt.Printf("%#v\n", strings.HasPrefix(strings.ToUpper("eos:111就斤斤计较"), strings.ToUpper("EOS")))

	persons := Persons{
		{
			Name: "test1",
			Age:  20,
		},
		{
			Name: "test2",
			Age:  22,
		},
		{
			Name: "test3",
			Age:  21,
		},
	}
	for i := 0; i < 100000; i++ {
		persons = append(persons, Person{Name: "test" + strconv.Itoa(i), Age: int64(i)})
	}
	fmt.Println("排序前")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Age)
	}
	sort.Sort(persons)
	fmt.Println("排序后")
	for _, person := range persons {
		fmt.Println(person.Name, ":", person.Age)
	}
	fmt.Println(1547 % 120)
	fmt.Println((1547 - 107) / 120)

	size := 1547
	step := size / 10
	tempSize := size - step*10
	var temp []int
	for i := tempSize; i < size; i += step {
		temp = append(temp, i)
	}
	if len(temp) > 0 {
		temp[len(temp)-1] = 888
	}
	fmt.Println(len(temp), temp)
	for i := 0; i < 100000; i++ {
		fmt.Println(strings.Contains(strings.ToUpper("Eosx"), strings.ToUpper("eos")))
	}
	v := strings.Split("eosbetdice11:8871964,", ",")
	fmt.Println(len(v[0]))

	WhiteList := "eosbetdice11:887196"
	var isTotals = true
	var total int64
	if strings.Contains(WhiteList,":"){
		split := strings.Split(WhiteList, ":")
		if len(split) >= 2{
			if split[0] == "eosbetdice11" {
				x ,_:= strconv.Atoi(split[1])
				total = int64(x)
				isTotals = false
			}
		}
	}
	fmt.Println(total,isTotals,time.Now())
}
