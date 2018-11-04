package main

import (
	"yhgolang/mkw/crawler/engine"
	"yhgolang/mkw/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})

}
