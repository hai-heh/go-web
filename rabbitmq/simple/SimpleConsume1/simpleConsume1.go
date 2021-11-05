package main

import rabbit "web/rabbitmq/init"

func main() {
	simple := rabbit.NewRabbitMQSimple("hello")
	simple.ConsumeSimple()
}