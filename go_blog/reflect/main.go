package main

import (
	"fmt"
	"reflect"
)

type User struct{
	Name string `json:"name" sql:"name"`
	Age int `json:"age"`
}

func main(){
	var u User
	u = User{
		Name:"jdx",
		Age:12,
	}
	t := reflect.TypeOf(u)
	fieldNum := t.NumField() // 2
	//获取结构体内字段
	for i := 0; i < fieldNum;i++{
		fmt.Printf("field:name:%v,type:%v,tag:%v\n",t.Field(i).Name,t.Field(i).Type,t.Field(i).Tag)
	}
	//field:name:Name,type:string,tag:json:"name" sql:"name"
	//field:name:Age,type:int,tag:json:"age"

	v := reflect.ValueOf(u)
	fmt.Println(v)
	fieldNum = v.NumField()
	for i := 0; i < fieldNum;i++{
		fmt.Println(v.FieldByName(v.Type().Field(i).Name))
	}
	//jdx
	//12

}
