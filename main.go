package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"shop.zozoo.net/config"
)

//初始化项目
func main() {
	r := gin.Default()

	//初始化配置文件
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	//设置静态目录
	r.Static("/static", "./static")
	//引入模板文件
	r.LoadHTMLGlob("template/*")

	//首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "gin Shop"})
	})
	//登陆页
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{"title": "gin Shop"})
	})
	//关于我们
	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{"title": "gin Shop"})
	})

	//关于我们
	r.GET("/details", func(c *gin.Context) {
		c.HTML(http.StatusOK, "details.html", gin.H{"title": "gin Shop"})
	})

	//购物车
	r.GET("/shopCart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "shopCart.html", gin.H{"title": "gin Shop"})
	})

	//母婴资讯
	r.GET("/information", func(c *gin.Context) {
		c.HTML(http.StatusOK, "information.html", gin.H{"title": "gin Shop"})
	})

	//购物车
	r.GET("/buyCart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "buyCart.html", gin.H{"title": "gin Shop"})
	})

	//商品分类
	r.GET("/commotidy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "commotidy.html", gin.H{"title": "gin Shop"})
	})

	r.Run(viper.GetString("server.port"))
}
