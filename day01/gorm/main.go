package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *sql.DB
)

func initDB() {
	dsn := "root:root@(127.0.0.1:3306)/test"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("sql open err %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("db ping err:%v\n", err)
		return
	}
	fmt.Println("数据库连接成功!")
}

func main() {
	initDB()
}
