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

	//获取分类
	var category []model.Category
	connect.Db.Select([]string{"id", "level","name","sort"}).Where("status=?",1).Find(&category)

	//定义变量，保存对数据资源进行排序
	var list map[string][]model.Cate = map[string][]model.Cate{}
	for _,cate := range category {
		if cate.Level == 0 {
			for _,item := range category{
				if item.Level != 0 && uint(item.Level) == cate.ID {
					var data model.Cate = model.Cate{
						item.ID,
						item.Name,
						item.Level,
						item.Sort,
					}
					list[cate.Name] = append(list[cate.Name], data)
				}
			}
		}
	}

	c.Set("dbAdv",map[string]interface{}{"adv":adv,"cate":&list})
}
