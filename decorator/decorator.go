package decorator

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	. "shop.zozoo.net/connect"
)

//缓存装饰器函数
func CacheDecorator(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//链接redis
		redisPool := RedisPool()
		redisPool.Close()

		conn := redisPool.Get()
		shopAdv,err := redis.String(conn.Do("get","shop_adv"))
		if err != nil {
			fmt.Println("读取出错")
		}else if shopAdv == "" {
			fmt.Println("缓存为空")
		}

		h(c)
	}
}
