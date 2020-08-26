package main

import (
	"fmt"
	"sync"
)
//竞争
var wg sync.WaitGroup
var v = 0
var m sync.Mutex
func f(){
	m.Lock()
	for i:=0;i<100000;i++ {
		v += 1
	}
	m.Unlock()
	defer wg.Done()
}

func main(){
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	fmt.Println(v)
}
