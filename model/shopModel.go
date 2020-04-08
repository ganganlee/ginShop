package model

import "github.com/jinzhu/gorm"

//商品表
type Shop struct {
	gorm.Model
	CategoryId uint                               //分类ID
	Name       string  `gorm:"size:256;not null"` //商品名称
	Price      float64 `gorm:"not null"`          //商品价格
	Thumbnail  string  `gorm:"size:125;not null"` //商品缩略图
	Poster     string  `gorm:"not null"`          //商品海报
	Size       string                             //商品尺寸
	Color      string                             //商品颜色
	Detail     string `gorm:"type:text"`          //商品介绍
	Status     uint8  `gorm:"default:1"`          //商品状态
	ShopCar    []ShopCar                          //定义一对多购物车模型关联
	Collection []Collection                       //定义一对多收藏模型关联
}
