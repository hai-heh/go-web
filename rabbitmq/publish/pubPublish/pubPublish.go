package main

import (
	"strconv"
	"time"
	rabbit "web/rabbitmq/init"
)

func main() {
	sub := rabbit.NewRabbitMQPubSub("exchange")
	for i:=0;i<100;i++{
		time.Sleep(time.Second*1);
		sub.PublishPub("hello"+strconv.Itoa(i))
	}
}
