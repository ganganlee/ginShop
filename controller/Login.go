package controller

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "shop.zozoo.net/common"
	. "shop.zozoo.net/connect"
	"shop.zozoo.net/model"
	"strconv"
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
		ResponseError(c,"","用户名不能为空")
		return
	}
	password := c.PostForm("username")
	if password == ""{
		ResponseError(c,"","用户密码不能为空")
		return
	}

	//实例化用户结构体
	userInfo := new(model.User)

	//加密密码
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))

	//获取用户数据
	Db.Where("username=? and password=?",username, password).First(userInfo)
	if userInfo.ID == 0 {
		ResponseError(c,gin.H{"username":username,"password":password},"用户名或密码错误")
		return
	}

	//将用户信息转为json字符串
	userByte,_ := json.Marshal(userInfo)
	//生成加密key
	keyByte := []byte("shop_user_"+strconv.Itoa(int(userInfo.ID)))
	userKey := fmt.Sprintf("%x", md5.Sum(keyByte))
	//获取cookie存入缓存，返回成功
	redisPool := RedisPool()
	defer redisPool.Close()
	conn := redisPool.Get()

	conn.Do("setex",userKey,24*60*60,string(userByte))

	//加入cookie
	//c.SetCookie("userInfo",userKey,24*60*60,"/","localhost", http.SameSiteNoneMode, true,true)
	c.SetCookie("userInfo", userKey, 3600*24, "/", "localhost", false, true)

	ResponseSuccess(c,"","ok")
}
