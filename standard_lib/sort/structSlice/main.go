package main

import (
	"fmt"
	"sort"
	"strconv"
)

//对[]struct 排序

type User struct{
	Name string
	Id int
}

type UserSlice []User

func (u UserSlice)Len() int {
	return len(u)
}

func (u UserSlice)Less(i,j int) bool{
	return u[i].Id > u[j].Id
}

func (u UserSlice)Swap(i,j int){
	u[i],u[j] = u[j],u[i]
}

func InitUserSlice(num int) UserSlice{
	//指定len>0 再append 会从num末尾加入，前面num个元素都是nil
	//[{ 0} { 0} { 0} { 0} { 0} {tom1 1} {tom2 2} {tom3 3} {tom4 4} {tom5 5}]
	//解决方案
	//1、var l = make(UserSlice,5) l[i] = u
	//2、var l = make(UserSlice,0) 或 var l = UserSlice ; l = append(l,u); append会自动make()
	var l = make(UserSlice,num)
	for i := 0; i < num; i++ {
		u := User{
			Name:fmt.Sprintf("tom%s",strconv.Itoa(i+1)),
			Id:i+1,
		}
		//l[i] = u
		l = append(l,u)
	}
	return l
}

func main(){
	fmt.Println(InitUserSlice(5))
	//sortStruct()
}

func sortStruct(){
	l := InitUserSlice(5)
	fmt.Println(l)
	sort.Sort(l)
	fmt.Println(l)
	sort.Sort(sort.Reverse(l))
	fmt.Println(l)
}