package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main(){
	//withCancel()
	//withDeadLine()
	//withTimeout()
	withValue()
	fmt.Println("done")
}

func withValue(){
	var wg sync.WaitGroup
	key := "kfuin"
	val := 1
	ctx := context.WithValue(context.Background(),key,val)
	wg.Add(1)
	go v1(ctx,&wg)
	wg.Add(1)
	go v2(ctx,&wg)
	wg.Wait()
}

func v1(ctx context.Context,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(ctx.Value("kfuin"))
}

func v2(ctx context.Context,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(ctx.Value("kfuin"))
}

func withTimeout(){
	var wg sync.WaitGroup
	ctx,_ := context.WithTimeout(context.Background(),time.Second*5)
	//但在任何情况下调用它的cancel函数都是很好的实践
	//defer cancel()

	wg.Add(1)
	go f2(ctx,&wg)

	wg.Wait()
}

func withDeadLine(){
	var wg sync.WaitGroup
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))
	// 尽管不主动调用cancel 也会达到超时通知子goroutine退出的效果。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	wg.Add(1)
	go f2(ctx,&wg)
	wg.Wait()
}

func withCancel(){

	ctx,cancel := context.WithCancel(context.Background())
	//通知子goroutine退出 即使子goroutine内部还有子goroutine仍然能够将其后代goroutine退出
	defer cancel()

	go f(ctx)

	time.Sleep(time.Second*5)

}

func f(ctx context.Context){
	LOOP:
	for{
		select{
		case <-ctx.Done()://等待通知
			break LOOP
			default:
		}
		fmt.Println("在吗？")
		time.Sleep(time.Second)
	}
}

func f2(ctx context.Context,wg *sync.WaitGroup){
	defer wg.Done()
LOOP:
	for{
		select{
		case <-ctx.Done()://等待通知
			break LOOP
		default:
		}
		fmt.Println("在吗？")
		time.Sleep(time.Second)
	}
}