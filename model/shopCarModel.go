package model

import "github.com/jinzhu/gorm"

//购物车数据表
type ShopCar struct {
	gorm.Model
	UserId    uint                              //用户ID
	ShopId    uint                              //商品ID
	Name      string `gorm:"size:256;not null"` //商品名称
	Thumbnail string `gorm:"size:125;not null"` //商品图片
	Size      string `gorm:"size:125"`          //商品尺寸
	Color     string `gorm:"size:125"`          //商品颜色
}
