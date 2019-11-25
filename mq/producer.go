package mq

import "github.com/streadway/amqp"

// Publish : 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel(RabbitURL) {
		return false
	}

	if nil == channel.Publish(
		exchange, // exchange
		routingKey,// routing key
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, //
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 消息持久化，虽然消息设置持久化了，但是并不能保证一定会
			ContentType:  "text/plain",
			Body:         msg}) {
		return true
	}
	return false
}
