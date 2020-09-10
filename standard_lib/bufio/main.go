package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	readFile()
}

func readFile(){
	f,err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(f)
	var s = make([]byte,1024)
	n,err := r.Read(s)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n",string(s))
	fmt.Println(n)
}
