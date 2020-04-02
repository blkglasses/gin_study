package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbAddress = "root:Xxf970524.@tcp(192.168.0.111:3306)/test"
var db *sqlx.DB

func initDB() (err error)  {
	db, err = sqlx.Connect("mysql", dbAddress)
	if err!=nil {
		fmt.Println("connect error")
		return
	}
	// 设置最大连接
	db.SetMaxOpenConns(100)
	// 最大空闲（连接池?)
	db.SetMaxIdleConns(20)

	driver := db.Driver()
	fmt.Printf("%#v",driver)
	return
}

func queryAllBook() (bookList []*Book, err error)  {
	sqlstr:="SELECT * FROM book"
	err = db.Select(&bookList, sqlstr)
	if err!=nil{
		fmt.Println("查询失败")
		return 
	}
	return 
}

func insertBook(title string, price int64) (err error)  {
	sqlstr:="INSERT INTO book(title,price) VALUES(?,?)"
	_, err = db.Exec(sqlstr, title, price)
	if err!=nil{
		fmt.Println("插入失败")
		return
	}
	return
}

func deleteBook(id int64) (err error) {
	sqlstr:="DELETE FROM book WHERE id=?"
	_, err = db.Exec(sqlstr, id)
	if err!=nil{
		fmt.Println("删除失败")
		return
	}
	return
}