package main

import rabbit "web/rabbitmq/init"

func main() {
	routing := rabbit.NewRabbitMQRouting("delayExchange", "delay")
	routing.PublishDelay("time out")
}
