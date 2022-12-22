package main

import "rabbitmq/Rabbitmq"

func main() {
	rabbitmq := Rabbitmq.NewRabbit("simple", "", "")
	rabbitmq.ConsumeSimple()
}
