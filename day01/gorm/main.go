package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *sql.DB
)

func initdb() {
	dsn := "root:123456@(127.0.0.1:3306)/db1"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("sql open err %v\n", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("db ping err:%v\n", err)
		return
	}
}

func main() {
	initdb()
}
