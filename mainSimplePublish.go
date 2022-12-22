package main

import (
	"fmt"
	"rabbitmq/Rabbitmq"
)

func main() {
	rabbitmq := Rabbitmq.NewRabbit("", "simple1", "")
	rabbitmq.PublishSimple("hello")
	fmt.Printf("发送成功")
}
