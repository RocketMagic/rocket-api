package util

import (
	"github.com/spf13/viper"
	"log"
)

// 获取配置
func GetConfig(path string) interface{} {
	viper.SetConfigName("config")  //设置配置文件的名字
	viper.AddConfigPath("./conf/") //添加配置文件所在的路径
	viper.SetConfigType("json")    //设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("config file error: %s\n", err)
	}

	return viper.Get(path)
}
