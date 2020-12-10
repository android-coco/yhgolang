package main

import yhrabbitmq "yhgolang/rabbitmq"

func main() {
	imooc := yhrabbitmq.NewRabbitMQTopic("exImoocTopic", "imooc.*.two")
	imooc.RecieveTopic()
}
