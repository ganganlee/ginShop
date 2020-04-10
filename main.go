package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"shop.zozoo.net/config"
	"shop.zozoo.net/connect"
	"shop.zozoo.net/route"
)

//初始化项目
func main() {
	r := gin.Default()

	//初始化配置文件
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	//初始化数据库连接
	err = connect.InitMysql()
	if err != nil {
		panic(err)
	}


	//controller.CreateAdv()
	//设置静态目录
	r.Static("/static", "./static")
	//引入模板文件
	r.LoadHTMLGlob("template/*")

	//加入路由
	route.CreateRoute(r)

	r.Run(viper.GetString("server.port"))
}
