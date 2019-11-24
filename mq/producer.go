package mq

import "github.com/streadway/amqp"

// Publish : 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel(RabbitURL) {
		return false
	}

	if nil == channel.Publish(
		exchange,
		routingKey,
		false, // 如果没有对应的queue, 就会丢弃这条消息
		false, //
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}