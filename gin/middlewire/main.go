package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//中间件使用


func main(){
	// 新建一个没有任何默认中间件的路由
	//r := gin.New()

	//也可以选择使用gin自带的中间件
	//r := gin.New()
	//r.Use(gin.Recovery())

	//默认注册 Logger 和 Recovery 中间件
	//Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
	//Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
	r := gin.Default()
	// 注册一个全局中间件
	r.Use(useTime())

	r.GET("/hand",func(c *gin.Context){
		num,_ := c.Get("num")
		//num2 := c.MustGet("num").(string) //如果取到的num不是string 直接panic
		c.JSON(200,gin.H{
			"msg":"ok",
			"num":num,
			//"num2":num2,
		})
	})

	//
	r.GET("/hand2",useTime(),func(c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"ok",
		})
	})

	r.Run()
}


//创建统计耗时的中间件
func useTime() gin.HandlerFunc {
	return func(c *gin.Context){
		start := time.Now().UnixNano()
		// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("num",1)

		//调用该请求的剩余处理程序
		c.Next()
		//不调用该请求的剩余处理程序
		//c.Abort()
		cost := time.Now().UnixNano() - start
		//应该记录到日志
		fmt.Printf("cost time:%d\n",cost)
	}
}
