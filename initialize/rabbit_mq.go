package initialize

import (
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/queue_jobs"
	"github.com/a20070322/shop-go/types"
	"github.com/streadway/amqp"
)

func RabbitMq() {
	global.RabbitMq = &types.RabbitMqType{}
	var err error
	configEnv := global.AppSetting.RabbitMq
	global.RabbitMq.Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s", configEnv.UserName, configEnv.PassWord, configEnv.IP, configEnv.Port, configEnv.VHost))
	if err != nil {
		fmt.Printf("RabbitMq 连接异常%s", err.Error())
		panic(err)
	}
	global.RabbitMq.PriorityCh, err = global.RabbitMq.Conn.Channel()
	if err != nil {
		fmt.Printf("RabbitMq PriorityCh 连接异常%s", err.Error())
		panic(err)
	}
	global.RabbitMq.DelayCh, err = global.RabbitMq.Conn.Channel()
	if err != nil {
		fmt.Printf("RabbitMq DelayCh 连接异常%s", err.Error())
		panic(err)
	}
	global.RabbitMq.OrderCh, err = global.RabbitMq.Conn.Channel()
	if err != nil {
		fmt.Printf("RabbitMq OrderCh 连接异常%s", err.Error())
		panic(err)
	}
	queue_jobs.DelayInit()
	queue_jobs.PriorityInit()
	queue_jobs.OrderInit()
	fmt.Println("RabbitMq初始化成功")

}
