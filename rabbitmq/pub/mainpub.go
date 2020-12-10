package main

import (
	"fmt"
	"strconv"
	"time"
	yhrabbitmq "yhgolang/rabbitmq"
)

func main() {
	rabbitmq := yhrabbitmq.NewRabbitMQPubSub("newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		fmt.Println("订阅模式生产第" +
			strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}
}
