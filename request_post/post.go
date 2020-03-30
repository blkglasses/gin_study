package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 1.创建路由
	// 默认使用2个中间件，Logger(), Recovery()
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/testPost", func(context *gin.Context) {
		name := context.PostForm("name")
		sex := context.PostForm("sex")
		var s = struct {
			name string
			sex  string
		}{}
		s.name = name
		s.sex = sex
		js,_:=json.Marshal(s)
		fmt.Println(s,js)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    s,
		})
	})

	// 可限制表单上传大小  默认32MB
	r.POST("/testFile", func(context *gin.Context) {
		file,_ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file, file.Filename)
		context.JSON(http.StatusOK, gin.H{
			"message": "上传成功",
		})
	})

	r.POST("/multipleFiles", func(c *gin.Context) {
		c.MultipartForm()
	})

	r.Run(":8000")
}
