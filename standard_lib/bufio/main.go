package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	readFile()
}

func readFile(){
	path,_:=filepath.Abs("./")
	f,err := os.Open(path + "/standard_lib/bufio/test.txt") //使用绝对路径
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

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
