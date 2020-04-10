package model

import "github.com/jinzhu/gorm"

//定义用户表
type User struct {
	gorm.Model
	Username string `gorm:"size:50;not null;unique" form:"username" json:"username" binding:"required,gte=4,lte=15"`           //用户名
	Password string `gorm:"size:255;not null" form:"password" json:"password" binding:"required,gte=6,lte=25,alphanum"` //用户密码
	Tel      string `gorm:"size:11;not null;unique" form:"tel" json:"tel" binding:"required,len=11"`                     //用户电话
	Avatar   string `gorm:"size:125"`                                                                                   //用户头像
	Ip       string `gorm:"size:20"`                                                                                    //注册ip
	Remarks  string `gorm:"size:255"`                                                                                   //备注
	Status   uint8                                                                                                      //状态
	//定义一对多收货地址模型关联
	Address []Address
	//定义一对多购物车模型关联
	ShopCar []ShopCar
	//定义一对多评论模型关联
	Comment []Comment
	//定义一对多订单模型关联
	Order []Order
	//定义一对多收藏模型关联
	Collection []Collection
}
