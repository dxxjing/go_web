package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{
	//binding:"required"  表必须字段 若为空 则报错
	Name string `form:"user" json:"user" binding:"required"`
	Password string  `form:"Password" json:"Password" binding:"required"`
}

func main() {

	g := gin.Default()

	getParam(g)

	g.Run()
}
//请求的参数通过URL路径传递，例如：/user/search/小王子/沙河
func pathParam(g *gin.Engine){
	g.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

}

func postParam(g *gin.Engine) {
	g.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		username := c.PostForm("username")
		address := c.DefaultPostForm("address","shanghai")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
}

//get 参数
func getParam(g *gin.Engine) {
	g.GET("user/search", func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		address := c.Query("address")
		//c.QueryArray("name")
		c.JSON(http.StatusOK, gin.H{
			"msg":     "ok",
			"name":    name,
			"address": address,
		})
	})
}

func bindParam(router *gin.Engine){
	// 绑定JSON的示例 ({"user": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var u User

		if err := c.ShouldBind(&u); err == nil {
			fmt.Printf("login info:%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"user":     u.Name,
				"password": u.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例 (user=q1mi&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var u User
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&u); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     u.Name,
				"password": u.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	router.GET("/loginForm", func(c *gin.Context) {
		var u User
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&u); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     u.Name,
				"password": u.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}
