package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()

	r.GET("/cookie",func(c *gin.Context){
		name,err := c.Cookie("name")
		if err != nil {
			c.SetCookie("name","jdx",3600,"/","localhost",false,true)
		}
		fmt.Println(name)
	})

	r.Run()
}
