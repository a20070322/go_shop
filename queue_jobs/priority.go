package queue_jobs

import (
	"encoding/json"
	"fmt"
	"github.com/a20070322/shop-go/global"

	"github.com/a20070322/shop-go/types"
	"github.com/streadway/amqp"
	"reflect"
)

const (
	PriorityQueueName      = "priority_queue"
	PriorityExchange   = "priority_exchange"
	PriorityRoutingKey = "priority_routingKey"
	PriorityConsume    = "priority_consume"
)

// 初始化公共队带有优先级
func PriorityInit() {
	var err error
	// 声明优先级队列
	_, err = global.RabbitMq.PriorityCh.QueueDeclare(PriorityQueueName, true, false, false, false,
		amqp.Table{
			//优先级定义
			"x-max-priority": 10,
		},
	)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 声明优先级队列交换机
	err = global.RabbitMq.PriorityCh.ExchangeDeclare(PriorityExchange, "direct", true, false, false, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 将延时队列处理队列绑定到延时队列交换机
	err = global.RabbitMq.PriorityCh.QueueBind(PriorityQueueName, PriorityRoutingKey, PriorityExchange, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	PriorityWorkOn()
}

// 发布延时任务队列
func PriorityPublish(payload *types.QueuePayloadType, priority int) error {
	// 序列化消息
	body, errJson := json.Marshal(payload)
	if errJson != nil {
		global.Logger.Error(errJson.Error())
		return errJson
	}
	err := global.RabbitMq.DelayCh.Publish(
		PriorityExchange,
		PriorityRoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			// 设置优先级
			Priority: uint8(priority),
			// 设置持久化存贮
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return err
}

// 消费任务
func PriorityWorkOn() {
	fmt.Println("公共优先级队列监听中")
	works, err := global.RabbitMq.DelayCh.Consume(
		PriorityQueueName,
		PriorityConsume,
		false,
		false,
		false,
		false,
		amqp.Table{
			"x-max-priority": int32(10),
		},
	)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	go func() {
		for d := range works {
			var payload types.QueuePayloadType
			errJson := json.Unmarshal(d.Body, &payload)
			if errJson != nil {
				global.Logger.Error("%s"+errJson.Error(), "PriorityWorkOn error")
				return
			}
			global.Logger.Info(fmt.Sprintf("公共优先级队列方法执行 [ %s ] : %s", payload.MethodName, payload.PayLoad))
			PriorityForwarding(d, &payload)
		}
	}()
}

// 任务调度器
func PriorityForwarding(d amqp.Delivery, payload *types.QueuePayloadType) {
	t := reflect.ValueOf(&PriorityQueue{})
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
