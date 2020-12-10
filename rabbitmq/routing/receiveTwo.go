package main

import yhrabbitmq "yhgolang/rabbitmq"

func main() {
	imoocTwo := yhrabbitmq.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocTwo.ReceiveRouting()
}
