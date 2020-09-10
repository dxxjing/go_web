package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main(){

	var str = "i am chinese"
	res := md5.Sum([]byte(str))
	fmt.Printf("%x\n",res)//6767731e6c67b5a05d71bbd7670e1ea0

	h := md5.New()
	io.WriteString(h, str)
	fmt.Printf("%x\n",h.Sum(nil))//6767731e6c67b5a05d71bbd7670e1ea0
}
