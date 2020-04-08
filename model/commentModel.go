package model

import "github.com/jinzhu/gorm"
//用户评论表
type Comment struct {
	gorm.Model
	UserId   uint   //用户id
	Username string //用户名称
	Avatar   string //用户头像
	ShopId   uint   //商品id
	Content  string //评论内容
}
