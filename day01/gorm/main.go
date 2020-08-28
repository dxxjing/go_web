package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *sql.DB //单例模式  连接池
)

type User struct{
	id int
	name string
	age int
}

func initDB()(err error){
	dsn := "root:root@(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	fmt.Println("数据库连接成功!")
	return
}

//单条查询
func queryOne(id int){
	sql := `select * from go_tbl where id = ?`
	var u User
	err := db.QueryRow(sql,id).Scan(&u.id,&u.name,&u.age)
	if err != nil{
		fmt.Printf("scan err:%v\n",err)
		return
	}
	fmt.Println(u)
}

func queryMore(){

	sql := `select * from go_tbl`
	rows,err := db.Query(sql)
	if err != nil {
		fmt.Printf("db query err:%v\n",err)
		return
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		if err = rows.Scan(&u.id,&u.name,&u.age);err != nil {
			fmt.Printf("scan err:%v\n",err)
			return
		}
		users = append(users,u)
	}
	fmt.Println(users)
}

func insert(name string,age int){
	sql := `insert into go_tbl (name,age) values (?,?)`
	res,err := db.Exec(sql,name,age)
	if err != nil {
		fmt.Println("exec err")
		return
	}
	id,_ := res.LastInsertId()
	fmt.Println("插入id:",id)
}

func update(id,age int){
	sql := `update go_tbl set age = ? where id = ?`
	res,err := db.Exec(sql,age,id)
	if err != nil{
		fmt.Printf("update err:%v\n",err)
		return
	}
	num,err := res.RowsAffected()
	if err != nil{
		fmt.Printf("get affected err:%v\n",err)
		return
	}
	fmt.Printf("影响行数%d\n",num)
}

func parpreQuery(){
	sql := `select * from go_tbl where id > ?`
	stmt,err := db.Prepare(sql)
	if err != nil {
		fmt.Println("prepare err",err)
		return
	}
	defer stmt.Close()

	rows,err := stmt.Query(0)
	if err != nil {
		fmt.Println("query err",err)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.id,&u.name,&u.age)
		users = append(users,u)
	}
	fmt.Println(users)
}

func prepareExec(){
	sql := `insert into go_tbl (name,age) values (?,?)`
	stmt,err := db.Prepare(sql)
	if err != nil {
		fmt.Println("prepare err",err)
		return
	}
	defer stmt.Close()


	if _,err := stmt.Exec("xiaomi",22);err != nil {
		fmt.Println("insert err",err)
		return
	}
	if _,err := stmt.Exec("dalong",33);err != nil {
		fmt.Println("insert err",err)
		return
	}

	fmt.Println("insert success")
}

func tx(){
	var (
		tx sql.Tx
		sql string
	)
	if tx,err := db.Begin();err != nil {
		fmt.Println("begin tx err",err)
		tx.Rollback()
		return
	}
	sql = `update go_tbl set age = 222 where id = ?`
	res,err := tx.Exec(sql,2)
	if err != nil {
		tx.Rollback()
		return
	}
	n1,err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return
	}

	sql = `update go_tbl set age = 222 where id = ?`
	res,err = tx.Exec(sql,3)
	if err != nil {
		tx.Rollback()
		return
	}
	n2,err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return
	}
	if err := tx.Commit(); err != nil{
		tx.Rollback()
		return
	}
	if n1 == 1 && n2 == 1 {
		fmt.Println("tx success")
	}else{
		fmt.Println("rollback")
	}

}



func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect db err:%v\n",err)
		return
	}
	//queryOne(1)
	//queryMore()

	prepareExec()
}


