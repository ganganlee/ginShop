package controller

import (
	"github.com/gin-gonic/gin"
	. "shop.zozoo.net/connect"
	"shop.zozoo.net/model"
)

func Index(c *gin.Context) {
	var adv = new(model.Advertisement)
	Db.Where("status=?",1).First(adv)
	c.Set("dbAdv",adv)
}
