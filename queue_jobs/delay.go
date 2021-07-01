package queue_jobs

import (
	"encoding/json"
	"fmt"
	"github.com/a20070322/shop-go/global"

	"github.com/a20070322/shop-go/types"
	"github.com/streadway/amqp"
	"reflect"
	"strconv"
)

const(
	// 死信队列
	DelayExpireWork = "delay_expire"
	// 任务处理队列
	DelayQueueWork = "delay_queue_work"
	// 交换机
	DelayExchangeName = "delay_exchange"
)

// 初始化延时队列
func DelayInit() {
	var err error
	// 声明延时队列
	_, err = global.RabbitMq.DelayCh.QueueDeclare(DelayExpireWork, true, false, false, false,
		amqp.Table{
			// 当消息过期时把消息发送到绑定的exchange
			"x-dead-letter-exchange": DelayExchangeName,
		},
	)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 声明延时队列处理队列
	_, err = global.RabbitMq.DelayCh.QueueDeclare(DelayQueueWork, true, false, false, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 声明延时队列转发交换机
	err = global.RabbitMq.DelayCh.ExchangeDeclare(DelayExchangeName, "fanout", true, false, false, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 将延时队列处理队列绑定到延时队列交换机
	err = global.RabbitMq.DelayCh.QueueBind(DelayQueueWork, "", DelayExchangeName, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	DelayWorkOn()

}

// 发布延时任务队列
func DelayPublish(payload *types.QueuePayloadType, expiration int) error {
	// 序列化消息
	body, errJson := json.Marshal(payload)
	if errJson != nil {
		global.Logger.Error(errJson.Error())
		return errJson
	}
	err := global.RabbitMq.DelayCh.Publish(
		"",
		DelayExpireWork,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			// 设置过期时间
			Expiration:  strconv.Itoa(expiration),
			// 设置持久化存贮
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return err
}

// 消费任务
func DelayWorkOn() {
	fmt.Println("延时队列监听中")
	works, err := global.RabbitMq.DelayCh.Consume(
		DelayQueueWork,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	go func() {
		for d := range works {
			var payload types.QueuePayloadType
			errJson := json.Unmarshal(d.Body, &payload)
			if errJson != nil {
				global.Logger.Error("%s"+errJson.Error(),"DelayWorkOn error")
				return
			}
			global.Logger.Info("延时队列方法执行 [ %s ] : %s",payload.MethodName,payload.PayLoad)
			DelayForwarding(d,&payload)
		}
	}()
}

// 任务调度器
func DelayForwarding(d amqp.Delivery, payload *types.QueuePayloadType) {
	t := reflect.ValueOf(&DelayQueue{})
	if t.MethodByName(payload.MethodName).IsZero() {
		global.Logger.Error("DelayForwarding  [%s] 方法不存在", payload.MethodName)
		d.Ack(true)
		return
	}
	t.MethodByName(payload.MethodName).
		Call([]reflect.Value{
			reflect.ValueOf(d),
			reflect.ValueOf(payload),
		})
}
