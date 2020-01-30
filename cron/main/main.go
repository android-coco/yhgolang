package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var x uintptr
	fmt.Println(x)
	fmt.Sprintf("1%s2", "fdsafasd")
	//https://xjwl-qp.oss-accelerate.aliyuncs.com/img/head_1.jpg
	//var wireteString = ""
	//for i := 1; i <= 2180; i++ {
	//	wireteString += strconv.Itoa(i) + "\n"
	//}
	//
	//var d1 = []byte(wireteString)
	//err2 := ioutil.WriteFile("./output2.txt", d1, 0666) //写入文件(字节数组)
	//check(err2)
}
