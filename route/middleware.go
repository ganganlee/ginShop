package route

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	. "shop.zozoo.net/connect"
	. "shop.zozoo.net/model"
)

//判断是否登陆中间件
func IsLoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//实例化用户表结构体
		userInfo := new(User)
		var isLogin bool
		//获取用户缓存key
		userKey, _ := c.Cookie("userInfo")
		fmt.Println(userKey)

		if userKey != "" {
			//前往redis获取数据，如果数据存在，说明还是登陆状态，否则未登录
			userKey = "shop_userinfo" + userKey

			redisPool := RedisPool()
			defer redisPool.Close()

			//获取数据
			conn := redisPool.Get()
			info, err := redis.Bytes(conn.Do("get", userKey))
			if err != nil {
				fmt.Println(err)
			}else {
				json.Unmarshal(info, userInfo)
				isLogin = true
			}
		}
		c.Set("userInfo", userInfo)
		c.Set("isLogin", isLogin)
		c.Next()
	}
}
