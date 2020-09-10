package main

import "fmt"

func main(){

	x := []int{2,3,5,7,11} //x 指向数组 {2,3,5,7,11} 其len = 5 cap = 5
	fmt.Println(x,len(x),cap(x))
	y := x[1:3] //y 指向数组 {3,5,7,11} 其 len = 2 cap = 5
	fmt.Println(y,len(y),cap(y))
	z := y[3:4] //z 指向数组 {11} 其len = 1 cap = 1
	fmt.Println(z,len(z),cap(z))
}
