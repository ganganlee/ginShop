package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Level  uint8 `gorm:"default:0"` //分类等级
	Name   string //分类名称
	Sort   uint //分类排序
	Status uint8 `gorm:"default:1"` //分类状态
}
