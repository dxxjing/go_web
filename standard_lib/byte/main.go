package main

import (
	"bytes"
	"fmt"
	"io"
)

func main(){
	//read()
	write()
}

func write(){
	var str string
	var res = make([]byte,64)
	buffer := bytes.NewBufferString(str)
	//将数据写入缓冲
	n,err := buffer.Write([]byte("ni hai hao ma!"))
	if err != nil{
		fmt.Println(err)
		return
	}
	//从缓冲读取出数据
	buffer.Read(res)
	fmt.Println(n,string(res))

}

func read(){

	var str = "ni hai hao ma"
	buffer := bytes.NewBuffer([]byte(str))

	var res = make([]byte,0)
	var p = make([]byte,5)
	for{
		n,err := buffer.Read(p)
		if err != nil{
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("已读取%d字节\n",n)
		res = append(res,p[:n]...)
	}
	fmt.Println(string(res))
}
