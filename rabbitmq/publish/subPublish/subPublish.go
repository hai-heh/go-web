package main

import rabbit "web/rabbitmq/init"

func main() {
	sub := rabbit.NewRabbitMQPubSub("exchange")
	sub.RecieveSub()
}
