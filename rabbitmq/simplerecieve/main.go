package main

import yhrabbitmq "yhgolang/rabbitmq"

func main() {
	rabbitmq := yhrabbitmq.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
}
