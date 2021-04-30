package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {})
	app.Get("/contact", func(ctx iris.Context) {
		_, _ = ctx.HTML("<h1> Hello from /contact </h1>")
	})
	app.Get("/test", before, mainHandler, func(ctx iris.Context) {
		value := ctx.FormValue("info")
		ctx.JSON(iris.Map{"hello": "iris MVC"})
		fmt.Println("test:", value)
	}, after)
	_ = app.Run(iris.Addr(":8888"))
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() //继续执行下一个handler，在本例中是mainHandler。
}
func after(ctx iris.Context) {
	println("After the mainHandler")
}
func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")
	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")
	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)
	ctx.Next() // 继续下一个handler 这里是after
}
