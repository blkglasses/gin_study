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
		js, _ := json.Marshal(s)
		fmt.Println(s, js)
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    s,
		})
	})

	// 可限制表单上传大小  默认32MB
	r.POST("/testFile", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		context.SaveUploadedFile(file, file.Filename)
		context.JSON(http.StatusOK, gin.H{
			"message": "上传成功",
		})
	})

	r.MaxMultipartMemory = 8 << 20
	r.POST("/multipleFiles", func(c *gin.Context) {
		// 获取表单
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		files := form.File["files"]
		// 遍历所有图片
		for _, f := range files {
			if err := c.SaveUploadedFile(f, f.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload success %d files", len(files)))
	})

	r.Run(":8000")
}
