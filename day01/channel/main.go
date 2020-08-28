package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup
//chan关闭 仍然可读
func main01() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		defer close(ch)

	}()
	time.Sleep(time.Second * 4)
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)
	x, ok := <-ch
	fmt.Println(x, ok)
	x, ok = <-ch
	fmt.Println(x, ok)
	x, ok = <-ch
	fmt.Println(x, ok)
	x, ok = <-ch
	//10 true
	//20 true
	//0 false
}
