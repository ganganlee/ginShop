package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop.zozoo.net/connect"
	"shop.zozoo.net/model"
)

func Index(c *gin.Context) {
	//获取广告缓存
	var adv = new(model.Advertisement)
	connect.Db.Where("status=?",1).First(adv)
	if adv.ID == 0 {
		fmt.Println("获取广告失败")
		adv = nil
	}
	c.Set("dbAdv",adv)
}
