package config

import (
	"github.com/spf13/viper"
	"os"
)

//读取配置文件
func InitConfig() error{
	//获取当前文件路径
	workDir,err := os.Getwd()
	if err != nil {
		return err
	}

	//设置配置文件名称
	viper.SetConfigName("config")

	//设置配置文件类型
	viper.SetConfigType("yaml")

	//设置配置文件路径
	viper.AddConfigPath(workDir+"\\config")

	//读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
