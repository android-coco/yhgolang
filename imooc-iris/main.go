package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"yhgolang/imooc-iris/web/controllers"
)

func main() {
	//1.创建实例
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//2.注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	//3.启动服务
	_ = app.Run(
		iris.Addr("localhost:8888"),
	)
}
