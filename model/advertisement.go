package model

//广告数据库
import "github.com/jinzhu/gorm"

type Advertisement struct {
	gorm.Model
	Poster string `gorm:"size:125;not null"` //广告海报
	Url    string `gorm:"size:125;not null"` //广告链接
	Status uint8  `gorm:"default:1"`         //广告状态
}
