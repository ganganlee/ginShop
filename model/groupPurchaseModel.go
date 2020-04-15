package model

import "github.com/jinzhu/gorm"

type GroupPurchase struct {
	gorm.Model
	ShopId   uint //商品id
	Status uint8 //状态
}
