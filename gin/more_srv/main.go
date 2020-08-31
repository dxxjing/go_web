package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

//同时开启多个服务
func main(){
	srv01 := &http.Server{
		Addr:":8080",
		Handler:route01(),
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}

	srv02 := &http.Server{
		Addr:":8081",
		Handler:route02(),
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		srv01.ListenAndServe()
	}()
	wg.Add(1)
	go func(){
		defer wg.Done()
		srv02.ListenAndServe()
	}()
	fmt.Println("server start....")
	wg.Wait()
}

func route01() http.Handler{
	r := gin.Default()

	r.GET("/test01",func(c *gin.Context){
		c.JSON(200,gin.H{
			"code":200,
			"msg":"test1",
		})
	})
	return r
}

func route02() http.Handler {
	r := gin.Default()
	r.GET("/test02",func(c *gin.Context){
		c.JSON(200,gin.H{
			"code":200,
			"msg":"test2",
		})
	})
	return r
}
