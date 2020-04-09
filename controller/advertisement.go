package controller

import (
	. "shop.zozoo.net/connect"
	. "shop.zozoo.net/model"
)

//添加海报广告
func CreateAdv() {
	//实例化广告结构体
	var adv Advertisement

	//连接数据库
	Db.Where("status=?",1).First(&adv)
}
