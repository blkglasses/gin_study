package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 1.创建路由
	// 默认使用2个中间件，Logger(), Recovery()
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world!")
	})
	r.GET("/username/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+action)
	})
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "xxf")
		c.String(http.StatusOK, name)
	})

	r.Run(":8000")
}