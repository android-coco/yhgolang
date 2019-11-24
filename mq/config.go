package mq

var RabbitURL = "amqp://youhao:123456@127.0.0.1:5672/"

const (
	AsyncTransferEnable = true

	TransExchangeName = "uploadserver.trans"

	TransOSSQueueName = "uploadserver.trans.oss"

	TransOSSErrQueueName = "uploadserver.trans.oss.err"

	TransOSSRoutingKey = "oss"
)
