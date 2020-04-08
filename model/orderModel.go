package model

import "github.com/jinzhu/gorm"

//订单数据表
type Order struct {
	gorm.Model
	UserId   uint
	Address  string `gorm:"size:255;not null"`
	Order    string `gorm:"size:125;not null"`  //订单号
	Status   uint8  `gorm:"default:0"`          //付款状态
	Ip       string `gorm:"size:125"`           //付款IP
	ShopInfo string `gorm:"type:text;nou null"` //包含的商品
	Remarks  string `gorm:"size:255"`           //备注
}
