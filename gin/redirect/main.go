package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	//http 重定向
	r.GET("/test",func(c *gin.Context){
		c.Redirect(301,"http://www.baidu.com")
	})
	//路由重定向
	r.GET("/test1",func(c *gin.Context){
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})

	r.GET("/test2",func(c *gin.Context){
		c.String(200,"test2")
	})

	r.Run()
}


