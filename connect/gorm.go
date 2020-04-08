package connect

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"shop.zozoo.net/model"
)

//初始化数据库链接
var Db *gorm.DB

func InitMysql() error {
	name := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	MaxIdleConns := viper.GetInt("mysql.MaxIdleConns")
	MaxOpenConns := viper.GetInt("mysql.MaxOpenConns")
	database := viper.GetString("mysql.database")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	db, err := gorm.Open("mysql", name+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	//设置连接池
	db.DB().SetMaxIdleConns(MaxIdleConns)
	db.DB().SetMaxOpenConns(MaxOpenConns)

	//数据迁移
	db.AutoMigrate(&model.User{},&model.Category{},&model.Address{},&model.Shop{},&model.ShopCar{},&model.Comment{},&model.Order{},&model.Collection{},&model.Advertisement{})

	Db = db
	return nil
}
