package initialize

import (
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/spf13/viper"
)

func Config() {
	v := viper.New()
	v.SetConfigFile("config.yaml") // 指定配置文件
	v.AddConfigPath("./")          // 指定查找配置文件的路径
	err := v.ReadInConfig()        // 读取配置信息
	if err != nil {                // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(&global.AppSetting); err != nil {
		fmt.Printf("err:%s", err)
	}
}
