package main

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"unsafe"
)

func main() {
	//var urlStr string = "http://api.1cloudsp.com/api/v2/send?accesskey=g5VyDRBvhb2J5jby&secret=ETCpd34PnWhholJxiuXl3NWT3XKmfxaQ&templateId=175515&mobile=13823682143&content=3211&sign=【便利卡】"
	escapeUrl := url.QueryEscape("【便利卡】")
	fmt.Println("编码:",escapeUrl)

	enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
	fmt.Println("解码:",enEscapeUrl)
	get, err := http.Get("http://api.1cloudsp.com/api/v2/send?accesskey=g5VyDRBvhb2J5jby&secret=ETCpd34PnWhholJxiuXl3NWT3XKmfxaQ&templateId=175515&mobile=13823682143&content=3211&sign=%E3%80%90%E4%BE%BF%E5%88%A9%E5%8D%A1%E3%80%91")
	fmt.Println(get,err)

	return

	var x struct {
		a bool
		b int16
		c []int
	}

	/**
	unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	*/

	/**
	uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	指针的运算
	*/
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"


	fmt.Println(reflect.TypeOf(x.b)) //uint

}
