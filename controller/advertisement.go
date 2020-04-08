package controller

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	. "shop.zozoo.net/connect"
	. "shop.zozoo.net/model"
)

//添加海报广告
func CreateAdv() {
	//获取缓存
	redisPool := RedisPool()
	defer redisPool.Close()
	conn := redisPool.Get()
	str, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(str)
		return
	}
	//实例化广告结构体
	var adv Advertisement

	adv.Poster = "/static/img/banner1.jpg"
	adv.Url = "http://www.baidu.com"
	//连接数据库
	Db.Create(&adv)

	if adv.ID == 0 {
		fmt.Println("创建失败")
	} else {
		fmt.Println("success")
	}
}
