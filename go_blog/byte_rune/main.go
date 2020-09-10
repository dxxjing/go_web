package main

import "fmt"

func main(){


	str := "ab你好"
	fmt.Println(len(str)) //8
	fmt.Println(len([]byte(str)))//8
	fmt.Println(len([]rune(str)))//4
}
