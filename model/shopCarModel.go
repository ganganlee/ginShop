package model

import "github.com/jinzhu/gorm"

//购物车数据表
type ShopCar struct {
	gorm.Model
	UserId uint
	ShopId uint
	Name   string //商品名称
	Img    string //商品图片
	Size   string //商品尺寸
	Color  string //商品颜色
}
