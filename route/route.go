package route

import (
	"github.com/gin-gonic/gin"
	"shop.zozoo.net/controller"
)

func CreateRoute(c *gin.Engine) {
	c.Use(IsLoginMiddleware())
	//用户注册
	c.POST("/register",controller.Register)

	//用户登陆
	c.GET("/login",controller.Login)
	c.POST("/login",controller.HandleLogin)
}
