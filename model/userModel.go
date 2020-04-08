package model

import "github.com/jinzhu/gorm"

//定义用户表
type User struct {
	gorm.Model
	Username   string `gorm:"size:50;not null"`        //用户名
	Password   string `gorm:"size:255;not null"`       //用户密码
	Tel        string `gorm:"size:11;not null;unique"` //用户电话
	Avatar     string `gorm:"size:125"`                //用户头像
	Ip         string `gorm:"size:20"`                 //注册ip
	Remarks    string `gorm:"size:255"`                //备注
	Address    []Address                               //定义一对多收货地址模型关联
	ShopCar    []ShopCar                               //定义一对多购物车模型关联
	Comment    []Comment                               //定义一对多评论模型关联
	Order      []Order                                 //定义一对多订单模型关联
	Collection []Collection                            //定义一对多收藏模型关联
}
