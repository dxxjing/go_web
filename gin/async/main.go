package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func main(){

	r := gin.Default()

	r.GET("/async",func(c *gin.Context){
		//注意：这里子goroutine必须使用gin.Context副本,否则会出问题
		var wg sync.WaitGroup
		tmpContext := c
		//异步执行
		wg.Add(1)
		go func(){
			defer wg.Done()
			time.Sleep(time.Second*3)
			fmt.Println("异步执行",tmpContext.Request.URL.Path)
		}()
		wg.Wait()
	})

	r.Run()
}
