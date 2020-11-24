package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	//将“before”处理程序注册为将要执行的第一个处理程序
	//在所有域的路由上。
	//或使用`UseGlobal`注册一个将跨子域触发的中间件。
	app.Use(before)

	//将“after”处理程序注册为将要执行的最后一个处理程序
	//在所有域的路由'处理程序之后。
	app.Done(after)

	// register our routes.
	app.Get("/", indexHandler)
	app.Get("/contact", contactHandler)

	_ = app.Run(iris.Addr(":8888"))
}

func before(ctx iris.Context) {
	// [...]
	fmt.Println("before")
	ctx.Next()
}
func after(ctx iris.Context) {
	// [...]
	fmt.Println("after")
}
func indexHandler(ctx iris.Context) {
	ctx.HTML("<h1>Index</h1>")
	ctx.Next() // 执行通过`Done`注册的“after”处理程序。
}
func contactHandler(ctx iris.Context) {
	// write something to the client as a response.
	ctx.HTML("<h1>Contact</h1>")
	ctx.Next() // 执行通过`Done`注册的“after”处理程序。
}
