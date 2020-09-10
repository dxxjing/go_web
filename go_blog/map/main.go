package main

import "fmt"

func main(){

	var s = make([]map[string]int,3)
	for index,val := range s{
		fmt.Println(index,val)
	}
	s[0] = make(map[string]int,10)
	s[0] = map[string]int{
		"a":1,
		"b":2,
	}
	for index,val := range s{
		fmt.Println(index,val)
	}

	var ss = make(map[string][]string,3)
	fmt.Println(ss)
	if _,ok := ss["jdx"];!ok{
		ss["jdx"] = make([]string,3)
	}
	ss["jdx"] = []string{"a","b"}
	fmt.Println(ss)
}
