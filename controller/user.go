package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "shop.zozoo.net/common"
	. "shop.zozoo.net/connect"
	. "shop.zozoo.net/model"
)

func Register(c *gin.Context) {
	isLogin := c.GetBool("isLogin")
	//判断是否登陆
	if isLogin {
		//已经登陆，重定向到个人中心页面
		c.Redirect(http.StatusTemporaryRedirect, "/home")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "gin Shop",
			"type":  "register", //login页面表单显示模式，login:登陆表单，register：注册表单
		})
	}
}
//用户注册
func HandlerRegister(c *gin.Context) {
	//获取用户注册信息
	var userInfo = new(User)
	err := c.ShouldBind(userInfo)
	if err != nil {
		ResponseError(c, "", err.Error())
		return
	}

	//判断用户名或者手机号是否在
	var user = new(User)
	Db.Select([]string{"id", "username", "tel"}).Where("username=?", userInfo.Username).Or("tel=?", userInfo.Tel).First(user)
	if user.ID != 0 {
		if user.Username == userInfo.Username {
			ResponseError(c, "", "此用户名已经被使用")
		} else {
			ResponseError(c, "", "此手机号已经被使用")
		}
		return
	}

	//创建数据

	//获取ip
	userInfo.Ip = c.ClientIP()

	//加密用户密码
	data := []byte(userInfo.Password)
	userInfo.Password = fmt.Sprintf("%x", md5.Sum(data))
	Db.Create(userInfo)

	//判断是否创建成功
	if userInfo.ID != 0 {
		ResponseSuccess(c, "", "ok")
	} else {
		ResponseError(c, "", "数据库繁忙")
	}
}
