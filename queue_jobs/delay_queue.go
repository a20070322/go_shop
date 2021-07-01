package queue_jobs

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/model/order"
	"github.com/a20070322/shop-go/types"
	"github.com/streadway/amqp"
)

type DelayQueue struct {
}

// 订单超时处理
func (del *DelayQueue) OrderTimeout(d amqp.Delivery, payload *types.QueuePayloadType) {
	boolOrder := order.OrderInfoInit(context.Background()).TimeOutOrder(payload.PayLoad)
	if boolOrder == false {
		global.Logger.Error("[%s] 订单c超时处理失败", payload.PayLoad)
	}
	d.Ack(true)
}

func (del *DelayQueue) Test(d amqp.Delivery, payload *types.QueuePayloadType) {
	fmt.Println(payload.PayLoad)
	fmt.Println("延时队列执行方法测试成功")
	d.Ack(true)
}
