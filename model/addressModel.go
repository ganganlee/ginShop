package model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	UserId  uint                              //用户ID
	Name    string `gorm:"size:125;not null"` //收件人名称
	Tel     string `gorm:"size:11;not null"`  //收件人电话
	Address string `gorm:"size:255;not null"` //收件人地址
	ZipCode uint16                            //收件人邮编
}
