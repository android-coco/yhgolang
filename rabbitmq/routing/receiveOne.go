package main

import yhrabbitmq "yhgolang/rabbitmq"

func main() {
	imoocOne := yhrabbitmq.NewRabbitMQRouting("exImooc", "imooc_one")
	imoocOne.ReceiveRouting()
}
