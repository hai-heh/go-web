package main

import rabbit "web/rabbitmq/init"

func main() {
	kutengone := rabbit.NewRabbitMQRouting("hehe", "bbb")
	kutengone.RecieveRouting()
}