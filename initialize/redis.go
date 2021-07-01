package initialize

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/go-redis/redis/v8"
)

func Redis() {
	ctx := context.Background()
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:    fmt.Sprintf("%s:%d",global.AppSetting.Redis.Host,global.AppSetting.Redis.Port),
		Password: global.AppSetting.Redis.Password, // no password set
		DB:       global.AppSetting.Redis.DB,          // use default DB
	})
	_, err := global.Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis is error")
		panic(err.Error())
	}
	fmt.Println("Redis初始化成功")
}