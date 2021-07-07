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
	OrderQueueName  = "order_queue"
	OrderExchange   = "order_exchange"
	OrderRoutingKey = "order_routingKey"
	OrderConsume    = "order_consume"
)

// 初始化订单处理业务队列，每次只取一个消息
func OrderInit() {
	var err error
	// 声明优先级队列
	_, err = global.RabbitMq.OrderCh.QueueDeclare(OrderQueueName, true, false, false, false,
		amqp.Table{
			//优先级定义
			"x-max-order": 10,
		},
	)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 声明优先级队列交换机
	err = global.RabbitMq.OrderCh.ExchangeDeclare(OrderExchange, "direct", true, false, false, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	// 将延时队列处理队列绑定到延时队列交换机
	err = global.RabbitMq.OrderCh.QueueBind(OrderQueueName, OrderRoutingKey, OrderExchange, false, nil)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	OrderWorkOn()
}

// 发布延时任务队列
func OrderPublish(payload *types.QueuePayloadType, order int) error {
	// 序列化消息
	body, errJson := json.Marshal(payload)
	if errJson != nil {
		global.Logger.Error(errJson.Error())
		return errJson
	}
	err := global.RabbitMq.DelayCh.Publish(
		OrderExchange,
		OrderRoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			// 设置优先级
			Priority: uint8(order),
			// 设置持久化存贮
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		global.Logger.Error(err.Error())
	}
	return err
}

// 消费任务
func OrderWorkOn() {
	fmt.Println("订单优先级队列监听中")
	works, err := global.RabbitMq.DelayCh.Consume(
		OrderQueueName,
		OrderConsume,
		false,
		false,
		false,
		false,
		amqp.Table{
			"x-max-order": int32(10),
		},
	)
	if err != nil {
		global.Logger.Panic(err.Error())
	}
	//设置每次从消息队列获取任务的数量
	if errQos := global.RabbitMq.DelayCh.Qos(1, 0, false); errQos != nil {
		global.Logger.Panic(errQos.Error())
	}
	go func() {
		for d := range works {
			var payload types.QueuePayloadType
			errJson := json.Unmarshal(d.Body, &payload)
			if errJson != nil {
				global.Logger.Error("%s"+errJson.Error(), "OrderWorkOn error")
				return
			}
			global.Logger.Info(fmt.Sprintf("order队列方法执行 [ %s ] : %s", payload.MethodName, payload.PayLoad))
			OrderForwarding(d, &payload)
		}
	}()
}

// 任务调度器
func OrderForwarding(d amqp.Delivery, payload *types.QueuePayloadType) {
	t := reflect.ValueOf(&OrderQueue{})
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
