package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接受数据的结构体
// binding:"required" 表示id为必传参数
type Student struct {
	Id   int    `form:"id" json:"id" xml:"id" binding:"required"`
	Name string `form:"username" json:"username" xml:"username"`
}

func main() {
	r := gin.Default()
	// json数据绑定结构体
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Student
		// 将request的body中的数据，自动绑定到json结构体中
		if err := c.ShouldBindJSON(&json); err != nil {
			// gin.H 封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 打印接受的数据
		fmt.Println(json)
	})

	// form表单数据绑定结构体
	r.POST("/loginForm", func(c *gin.Context) {
		var form Student
		// Bind() 默认绑定form数据
		if err := c.ShouldBind(&form); err != nil {
			// gin.H 封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("%#v", form)
	})

	// uri数据绑定结构体
	r.GET("/loginUri/:Id/:Name", func(c *gin.Context) {
		var form Student
		// Bind() 默认绑定form数据
		if err := c.ShouldBindUri(&form); err != nil {
			// gin.H 封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("%#v", form)
	})
	r.Run(":8000")
}
