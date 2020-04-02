package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	err := initDB()
	if err!=nil{
		panic(err)
	}
	r:=gin.Default()
	// 加载页面
	r.LoadHTMLGlob("./book_mgr/templates/*")
	// 查询所有图书
	r.GET("/book/list",bookListHandler)
	r.Run(":8000")
}
func bookListHandler(c *gin.Context)  {
	bookList, err := queryAllBook()
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":err})
	}
	c.HTML(http.StatusOK,"book_list.html", gin.H{"data": bookList})
}