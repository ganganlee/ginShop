package route

import (
	"github.com/gin-gonic/gin"
	"shop.zozoo.net/controller"
	. "shop.zozoo.net/decorator"
)

func CreateRoute(c *gin.Engine) {
	c.Use(IsLoginMiddleware())
	//用户注册
	c.GET("/register",controller.Register)
	c.POST("/register",controller.HandlerRegister)

	//用户登陆
	c.GET("/login",controller.Login)
	c.POST("/login",controller.HandleLogin)

	//商城首页
	c.GET("/",CacheDecorator(controller.Index))
}
