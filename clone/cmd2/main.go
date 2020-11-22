package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	//各种中间件
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())
	//静态文件
	router.StaticFS("/static", http.Dir("../static"))
	//router.StaticFS("/mnt", http.Dir("./mnt"))
	//router.StaticFile("/favicon.ico", "../static/logo.png")
	//router.LoadHTMLGlob("../view/**/*")
	router.LoadHTMLGlob("../view/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run(":9090")
}
