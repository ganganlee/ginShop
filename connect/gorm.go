package connect

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

//初始化数据库链接
var Db *gorm.DB
func InitMysql()error{
	name := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	MaxIdleConns:= viper.GetInt("mysql.MaxIdleConns")
	MaxOpenConns:= viper.GetInt("mysql.MaxOpenConns")
	db, err := gorm.Open("mysql", name+":"+password+"@/gin?charset="+charset+"&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(MaxIdleConns)
	db.DB().SetMaxOpenConns(MaxOpenConns)

	//数据迁移
	//db.AutoMigrate(&model.User{})
	Db = db
	return nil
}
