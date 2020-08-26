package main

import "fmt"

func main(){
	//var a *int  未初始化 nil
	var a = new(int)
	*a = 100
	fmt.Println(*a)
}
