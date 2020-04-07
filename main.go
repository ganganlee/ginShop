package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//初始化项目
func main(){
	r:= gin.Default()

	//设置静态目录
	r.Static("/static", "./static")
	//引入模板文件
	r.LoadHTMLGlob("template/*")

	//首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{"title":"gin Shop"})
	})
	//登陆页
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",gin.H{"title":"gin Shop"})
	})
	//关于我们
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK,"about.html",gin.H{"title":"gin Shop"})
	})
	//关于我们
	r.GET("/details", func(c *gin.Context) {
		c.HTML(http.StatusOK,"details.html",gin.H{"title":"gin Shop"})
	})
	r.Run()
}
