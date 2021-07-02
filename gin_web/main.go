package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/*")
	router.GET("/test/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "weixin://dl/business/?t=fYeZEvDz1hq",
		})
	})
	// 1.异步
	router.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			copyContext.JSON(200, gin.H{"1": "2"})
			log.Println("异步执行：" + copyContext.Request.URL.Path,copyContext)
		}()
	})
	// 2.同步
	router.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
		c.JSON(200, gin.H{"1": "2"})
	})
	router.Run(":9999")
}
