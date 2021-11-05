package main

import (
	"strconv"
	rabbit "web/rabbitmq/init"
)

func main() {
	pub1 := rabbit.NewRabbitMQRouting("hehe", "aaa")
	pub2 := rabbit.NewRabbitMQRouting("hehe", "bbb")
	for i:=0;i<=1;i++ {
		pub1.PublishRouting("hello"+strconv.Itoa(i))
		pub2.PublishRouting("world"+strconv.Itoa(i))
	}
}