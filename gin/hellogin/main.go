package main

import "github.com/gin-gonic/gin"

func main(){
	//创建默认路由引擎
	g := gin.Default()

	//绑定路由规则 执行函数
	//gin.Context 封装了 Request  Response
	g.GET("/hello",func(ctx *gin.Context){
		ctx.String(200,"hello gin")
	})
	//监听端口  默认8080
	g.Run(":8081")
}
