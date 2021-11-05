package main

import (
	"strconv"
	"time"
	rabbit "web/rabbitmq/init"
)

func main() {
	simple := rabbit.NewRabbitMQSimple("hello")
	for i:=0;i<100;i++{
		time.Sleep(1*time.Second)
		simple.PublishSimple("测试"+strconv.Itoa(i))
	}
}
