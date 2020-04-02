package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	// 加载页面
	r.LoadHTMLGlob("./book_mgr/templates/*")
	// 查询所有图书
	r.GET("/book/list", bookListHandler)
	r.GET("/book/new", openNewPageHandler)
	r.POST("/book/new", insertBookHandler)
	r.GET("/book/delete",deleteBookHandler)
	r.Run(":8000")
}
func bookListHandler(c *gin.Context) {
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.HTML(http.StatusOK, "book_list.html", gin.H{"data": bookList})
}

func openNewPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", gin.H{})
}

func insertBookHandler(c *gin.Context) {
	title := c.PostForm("title")
	p := c.PostForm("price")
	price, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数（价格）类型错误"})
		return
	}
	err = insertBook(title, int64(price))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "新增数据失败"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"新增成功"})
}

func deleteBookHandler(c *gin.Context)  {
	keyId := c.Query("id")
	id, err := strconv.Atoi(keyId)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"请求参数类型错误"})
		return
	}
	err = deleteBook(int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "删除数据失败"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"删除成功"})
}