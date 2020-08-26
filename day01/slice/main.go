package main

import "fmt"

func main(){
	var s = [...]int{1,2,3,4,5,6,7,8}
	s1 := s[:]
	s2 := append(s1,9)
	fmt.Println(s1)
	fmt.Println(s2)
}
