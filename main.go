package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	ginServer := gin.Default()

	ginServer.LoadHTMLGlob("template/*")
	ginServer.Static("/static", "./static")

	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "go 后台",
		})
	})

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})

	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "post user"})
	})

	//restful 风格
	ginServer.GET("/user/info/:username/:age", func(context *gin.Context) {
		username := context.Param("username")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{"username": username, "age": age})
	})

	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	ginServer.Run(":8082")

}
