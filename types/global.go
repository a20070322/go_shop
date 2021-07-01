package types

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMqType struct {
	Conn       *amqp.Connection `json:"conn"`
	PriorityCh *amqp.Channel    `json:"priority_ch"`
	DelayCh    *amqp.Channel    `json:"delay_ch"`
	OrderCh    *amqp.Channel    `json:"order_ch"`
}

func (r RabbitMqType) Close() {
	if err := r.PriorityCh.Close(); err != nil {
		log.Println("RabbitMq [PriorityCh] 通道关闭异常", err)
	}
	if err := r.DelayCh.Close(); err != nil {
		log.Println("RabbitMq [DelayCh] 通道关闭异常", err)
	}
	if err := r.Conn.Close(); err != nil {
		log.Println("RabbitMq [Conn] 连接关闭异常", err)
	}
	log.Println("RabbitMq stop success")
}

type QueuePayloadType struct {
	MethodName string `json:"method_name"`
	PayLoad    string `json:"pay_load"`
}
