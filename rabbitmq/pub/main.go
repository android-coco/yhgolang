package main

import yhrabbitmq "yhgolang/rabbitmq"

func main() {
	rabbitmq := yhrabbitmq.NewRabbitMQPubSub("newProduct")
	rabbitmq.ReceiveSub()
}
