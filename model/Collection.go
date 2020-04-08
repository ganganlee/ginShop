package model

import "github.com/jinzhu/gorm"

//收藏数据表
type Collection struct {
	gorm.Model
	UserId    uint
	ShopId    uint
	Name      string `gorm:"size:256;not null"` //商品名称
	Thumbnail string `gorm:"size:125;not null"` //商品缩略图
}
