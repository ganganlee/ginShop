package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "shop.zozoo.net/model"
)

//用户注册
func Register(c *gin.Context) {
	//获取用户注册信息
	var userInfo = new(User)
	err := c.ShouldBind(userInfo)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(userInfo)
	}
}
