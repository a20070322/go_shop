package queue_jobs

import (
	"fmt"
	"github.com/a20070322/shop-go/types"
	"github.com/streadway/amqp"
	"time"
)

type PriorityQueue struct {

}
func (del *PriorityQueue) Test(d amqp.Delivery, payload *types.QueuePayloadType) {
	time.Sleep(1 * time.Second)
	fmt.Println(payload.PayLoad)
	fmt.Println("优先级方法测试")
	d.Ack(true)
}
