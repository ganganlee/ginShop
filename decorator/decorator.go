package decorator

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	. "shop.zozoo.net/connect"
	"shop.zozoo.net/model"
)

//缓存装饰器函数
func CacheDecorator(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//链接redis
		redisPool := RedisPool()
		defer redisPool.Close()

		conn := redisPool.Get()
		shopAdv, err := redis.Bytes(conn.Do("get", "shop_adv"))
		//判断缓存出错或者缓存不存在
		shopAdv = nil
		if err != nil || shopAdv == nil {
			//查询数据库
			h(c)
			res, exists := c.Get("dbAdv")
			if !exists {
				//数据库不存在
				res = nil
			}
			shopAdv, _ = json.Marshal(res)

			conn.Do("setex", "shop_adv", 24*60*60, string(shopAdv))
		}
		//将字符串解析为结构体输出
		var data map[string]interface{} = map[string]interface{}{
			"adv":  new(model.Advertisement),
			"cate": map[string][]model.Cate{},
		}

		err = json.Unmarshal(shopAdv, &data)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(data)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "商城首页",
			"adv":   data["adv"],
			"cate":  data["cate"],
		})
	}
}
