package main

import rabbit "web/rabbitmq/init"

func main() {
	mq := rabbit.NewRabbitMQRouting("delayExchange", "delay")
	mq.RecieveDelay()
}
