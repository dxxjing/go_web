package main

import "fmt"

type User struct{
	Name string
	Age int
}

func main(){

	var a1 [5]int
	a1 = [5]int{1,2,3,4,5}
	fmt.Println(a1)

	var a3 [5]int = [5]int{1,2,3,4,5}
	fmt.Println(a3)

	a2 := [...]string{"a","b","c"}
	fmt.Println(a2)

	a4 := [...]User{
		{Name:"jdx",Age:13},
		{Name:"tom",Age:15},
	}
	fmt.Println(a4)
}
