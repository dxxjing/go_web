package main

import (
	"encoding/json"
	"fmt"
)

//Person 人类
type Person struct {
	Name string `json:"name"` //名字
	Age  int    `json:"age"`  //年龄
}

func main() {
	var (
		p   Person
		ret []byte
		err error
	)

	p = Person{
		Name: "jdx",
		Age:  14,
	}

	ret, err = json.Marshal(p)
	if err != nil {
		fmt.Println("marshal err")
		return
	}
	fmt.Println(string(ret))

	json.Unmarshal(ret, &p)
	fmt.Println(p)
}
