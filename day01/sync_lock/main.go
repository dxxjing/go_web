package main

import (
	"fmt"
	"sync"
)
//竞争
var wg sync.WaitGroup
var v = 0

func f(){
	for i:=0;i<5000;i++ {
		v += i
	}
	defer wg.Done()
}

func main(){
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	fmt.Println(v)
}
