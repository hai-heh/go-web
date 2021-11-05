package main

import rabbit "web/rabbitmq/init"

func main() {
	kutengone := rabbit.NewRabbitMQRouting("example", "routingkey")
	kutengone.CanalConsumer()
}
