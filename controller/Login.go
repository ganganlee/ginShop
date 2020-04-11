package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	isLogin := c.GetBool("isLogin")
	//判断是否登陆
	if isLogin {
		//已经登陆，重定向到个人中心页面
		c.Redirect(http.StatusTemporaryRedirect, "/home")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "gin Shop",
			"type":  "login", //login页面表单显示模式，login:登陆表单，register：注册表单
		})
	}
}

func HandleLogin(c *gin.Context){
	//获取用户名与密码
	username := c.PostForm("username")
	if username == ""{
		common.ResponseError(c,"","用户名不能为空")
		return
	}
	password := c.PostForm("username")
	if password == ""{
		common.ResponseError(c,"","用户密码不能为空")
		return
	}

	//实例化用户结构体
	userInfo := new(model.User)

	//加密密码
	data := []byte(password)
	userInfo.Password = fmt.Sprintf("%x", md5.Sum(data))
}
