package model

import "github.com/jinzhu/gorm"
//用户评论表
type Comment struct {
	gorm.Model
	UserId   uint   //用户id
	Username string `gorm:"size:50;not null"`//用户名称
	Avatar   string `gorm:"size:125"` //用户头像
	ShopId   uint   //商品id
	Content  string `gorm:"size:255;not null"` //评论内容
}
