package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	//readFile()
	writeFile()
}

func writeFile(){
	f,err := os.OpenFile("standard_lib/bufio/write.txt",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.Write([]byte("ni hao"))
	w.Write([]byte("hello"))

	//带缓冲的写 最后必须调用flush 刷入文件
	w.Flush()
}

/**
	利用缓冲 每次读取5字节
 */
func readFile(){
	//默认从 GOPATH 目录下找文件
	//path,_:=filepath.Abs("./")
	//f,err := os.Open(path+"standard_lib/bufio/read.txt") //使用绝对路径
	f,err := os.Open("standard_lib/bufio/read.txt")
	if err != nil {
		fmt.Printf("%v\n",err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	var s = make([]byte,5)
	var res string
	for{
		n,err := r.Read(s)
		if err != nil{
			if err == io.EOF{
				fmt.Println("读取结束")
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Println("已读取",n,"字节")
		res += string(s)
	}
	fmt.Println(res)
}
