package main

import (
	"log"
	"sync"
)
var m sync.Map

func main(){
	// 建立对象
	var pipe = &sync.Pool{
		New:func()interface{}{
			return "Hello, BeiJing"
		},
	}

	// 准备放入的字符串
	val := "Hello,World!"

	// 放入
	pipe.Put(val)

	// 取出
	log.Println(pipe.Get())

	// 再取就没有了,会自动调用NEW
	log.Println(pipe.Get())
}