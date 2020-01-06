package main

import (
	"fmt"
	protoUtil "github.com/golang/protobuf/proto"
	"yhgolang/proto"
)

func main() {
	var attrs = make(map[string]*proto.Result)
	attrs["1"] = &proto.Result{
		Url:   "www.baidu.com",
		Title: "百度",
	}
	attrs["2"] = &proto.Result{
		Url:   "www.baidu.com",
		Title: "百度",
	}
	attrs["3"] = &proto.Result{
		Url:   "www.baidu.com",
		Title: "百度",
	}
	// 初始化消息
	score_info := &proto.Product{}
	score_info.Name = "1"
	score_info.Attrs = attrs

	// 以字符串的形式打印消息
	fmt.Println(score_info.String())

	// encode, 转换成二进制数据
	data, err := protoUtil.Marshal(score_info)
	if err != nil {
		panic(err)
	}

	// decode, 将二进制数据转换成struct对象
	new_score_info := proto.Product{}
	err = protoUtil.Unmarshal(data, &new_score_info)
	if err != nil {
		panic(err)
	}

	// 以字符串的形式打印消息
	fmt.Println(new_score_info.Attrs["1"].Url)

	var y []int
	X(&y)
	fmt.Println(y)
}

func X(x *[]int) {
	for i := 0; i < 10; i++ {
		*x = append(*x, i)
	}
}
