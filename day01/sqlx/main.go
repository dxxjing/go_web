package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

type User struct {
	ID   int	`"sql:id"`
	Name string	`"sql:name"`
	Age  int	`"sql:age"`
}

func initDB() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/test"
	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		return
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	fmt.Println("sqlx connect success")
	return
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println("sqlx connect err", err)
		return
	}
	//queryOne(1)
	//queryMore()
	insert()
}

func queryOne(id int) {
	sql := `select id,name,age from go_tbl where id = ?`
	var u User
	if err := db.Get(&u, sql, id); err != nil {
		fmt.Println("get data err", err)
		return
	}
	fmt.Printf("%#v\n",u)
}

func queryMore(){
	sql := `select * from go_tbl`
	var users []User
	//注意：这里users必须传入地址 因为go所有的类型都是值传递
	//但是为啥函数参数传入切片能改变原来的值？因为切片内部有个指针指向一个数组，改变其值就会改变其内部指向的数据，
	if err := db.Select(&users,sql);err != nil{
		fmt.Printf("select failed %v\n",err)
		return
	}
	fmt.Println(users)
}

//删除 更新 插入和原生sql一致
func insert(){
	sql := `insert into go_tbl (name,age) values ("iii",99)`
	res,err := db.Exec(sql)
	if err != nil{
		fmt.Println("insert err",err)
		return
	}
	id,_ := res.LastInsertId()
	fmt.Println("插入",id)
}
