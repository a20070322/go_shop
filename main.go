package main

import (
	"github.com/a20070322/shop-go/cmd"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/initialize"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func init() {
	//初始化配置文件
	initialize.Config()
	//初始化日志文件
	initialize.Logger(global.AppSetting.Logger.LogPath, global.AppSetting.Logger.LogErrPath, global.AppSetting.Logger.Level)
}

func main() {
	// fmt.Printf("pid: %dn", os.Getpid())
	//fmt.Println("当前版本：")
	//initialize.Router()
	app := &cli.App{
		Name:    "go-shop",
		Usage:   "go-shop服务端",
		Version: global.AppSetting.Version,
		//Flags: []cli.Flag{
		//	&cli.StringFlag{
		//		Name:  "lang",
		//		Value: "english",
		//		Usage: "language for the greeting",
		//	},
		//},
		Action: func(context *cli.Context) error {
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
			//启动http服务
			initialize.Router()
			return nil
		},
		Commands: []*cli.Command{
			cmd.EntMigrate,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
