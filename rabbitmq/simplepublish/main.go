package main

import (
	"fmt"
	"strconv"
	"time"
	yhrabbitmq "yhgolang/rabbitmq"
)

func main() {
	rabbitmq := yhrabbitmq.NewRabbitMQSimple("imoocSimple")
	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Hello imooc!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("发送成功！", i)
	}
}
