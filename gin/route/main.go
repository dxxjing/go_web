package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	g := gin.Default()

	//initRoute(g)
	anyRoute(g)

	noRoute(g)
	g.Run()
}

//匹配所有请求方法
func anyRoute(g *gin.Engine){
	g.Any("/test",func(c *gin.Context){
		c.String(200,"can accept get/post/put/delete")
	})
}
//找不到路由
func noRoute(g *gin.Engine){
	g.NoRoute(func(c *gin.Context){
		c.String(404,"not found")
	})
}

func initRoute(g *gin.Engine){
	g.GET("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"msg":"get",
		})
	})

	g.POST("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"msg":"post",
		})
	})

	g.PUT("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"msg":"put",
		})
	})

	g.DELETE("/book",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"msg":"delete",
		})
	})
}
