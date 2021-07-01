package queue_jobs

import (
	"context"
	"fmt"
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/model/order"
	"github.com/a20070322/shop-go/types"
	"github.com/streadway/amqp"
	"time"
)

type PriorityQueue struct {
}

// 订单处理
func (del *PriorityQueue) OrderProcessing(d amqp.Delivery, payload *types.QueuePayloadType) {
	boolOrder := order.OrderInfoInit(context.Background()).HandleOrder(payload.PayLoad)
	if boolOrder == false {
		global.Logger.Error("[%s] 订单处理失败", payload.PayLoad)
	}else{
		// 订单超时处理，超过20分钟未支付 取消订单且回复库存
		DelayPublish(&types.QueuePayloadType{
			MethodName: "OrderTimeout",
			PayLoad:    payload.PayLoad,
		}, 20*60*1000)
	}
	d.Ack(true)
}

func (del *PriorityQueue) Test(d amqp.Delivery, payload *types.QueuePayloadType) {
	time.Sleep(1 * time.Second)
	fmt.Println(payload.PayLoad)
	fmt.Println("优先级方法测试")
	d.Ack(true)
}
