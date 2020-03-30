package main

import "github.com/gin-gonic/gin"

func main() {
	r:=gin.Default()
	g1:=r.Group("/group1")
	// {} 是书写规范 不是语法  省略不报错
	{
		g1.GET("/login")
		g1.GET("/register")
	}
	g2:=r.Group("/group2")
	{
		g2.POST("/login")
		g2.POST("/register")
	}
	r.Run("8000")
}
