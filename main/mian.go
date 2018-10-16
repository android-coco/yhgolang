package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
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
const emailReg  = `^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$`

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
	now := time.Now()
	next := now.Add(time.Hour * 10)
	fmt.Println("xxxxx:" ,next)
	next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
	end  := time.Date(next.Year(), next.Month(), next.Day(), 23, 59, 59, 0, next.Location())

	fmt.Println(next.Format("2006-01-02 15:04:05"),end.Format("2006-01-02 15:04:05"))
	jsonstr := `{"status":"executed","cpu_usage_us":2841,"net_usage_words":19,"trx":{"id":"685550f6d82f30265087df37eebb2ed92fcee3276b6fd6895f2296d5a6303858","signatures":["SIG_K1_K6yQ7sRh5hvDU9LjfKfPQv572Eb9tcWNmSvC7ezfosMf8uf7i7LhyaijLwdLWkHoKR6w9ZsFiuTEc2E3jiGwGMBn9XtDin"],"compression":"none","packed_context_free_data":"","context_free_data":[],"packed_trx":"2d3dc45bc1e603dab9a3000000000100a6823403ea3055000000572d3ccdcd01a09865f94e93876600000000a8ed32323aa09865f94e938766401dce8db90931556e0000000000000004454f5300000000196d61743a3333313939393a313a6a676e6a676e6a676e31323300","transaction":{"expiration":"2018-10-15T07:09:33","ref_block_num":59073,"ref_block_prefix":2746866179,"max_net_usage_words":0,"max_cpu_usage_ms":0,"delay_sec":0,"context_free_actions":[],"actions":[{"account":"eosio.token","name":"transfer","authorization":[{"actor":"gu3tanrtgqge","permission":"active"}],"data":{"from":"gu3tanrtgqge","to":"eosknightsio","quantity":"0.0110 EOS","memo":"mat:331999:1:jgnjgnjgn123"},"hex_data":"a09865f94e938766401dce8db90931556e0000000000000004454f5300000000196d61743a3333313939393a313a6a676e6a676e6a676e313233"}],"transaction_extensions":[]}}}`
	fmt.Println(json.Unmarshal([]byte(jsonstr),&Person{}))

	matched, _ := regexp.MatchString(emailReg, "1111")
	fmt.Println(matched)
}


