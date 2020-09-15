package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func main(){

	r := gin.Default()
	//异步执行 无需waitgroup
	r.GET("/async",func(c *gin.Context){
		//注意：这里子goroutine必须使用gin.Context副本,否则会出问题 例如在协程内c.Set("name","小明") 由于协程未返回 此时在协程外是无法获取到
		//var wg sync.WaitGroup
		tmpContext := c.Copy()
		//异步执行
		fmt.Println("异步执行start....")
		//wg.Add(1)
		go func(){
			//defer wg.Done()
			time.Sleep(time.Second*5)
			fmt.Println("异步执行end",tmpContext.Request.URL.Path)
		}()
		//wg.Wait()
		fmt.Println("继续执行其他....")
		//输出：
		//异步执行start....
		//继续执行其他....
		//异步执行end
	})
	r.GET("/async",func(c *gin.Context){
		//注意：这里子goroutine必须使用gin.Context副本,否则会出问题
		//var wg sync.WaitGroup
		tmpContext := c
		//同步执行
		fmt.Println("同步执行start....")
		
		time.Sleep(time.Second*5)
		fmt.Println("同步执行end",tmpContext.Request.URL.Path)
		
		fmt.Println("继续执行其他....")
		//输出：
		//同步执行start....
		//同步执行end
		//继续执行其他....
		
	})

	r.Run()
}
