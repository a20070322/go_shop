package main

import (
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/initialize"
)

func init() {
	//初始化配置文件
	initialize.Config()
	//初始化日志文件
	initialize.Logger(global.AppSetting.Logger.LogPath, global.AppSetting.Logger.LogErrPath, global.AppSetting.Logger.Level)
	//	初始化redis
	initialize.Redis()
	//初始化ent
	initialize.Ent()
	//初始化微信小程序
	initialize.MiniChat()
	//初始化RabbitMq
	initialize.RabbitMq()
	//初始化微信支付
	initialize.WechatPay()
}

func main() {
	// fmt.Printf("pid: %dn", os.Getpid())


	//fmt.Println("当前版本：")
	initialize.Router()
}

