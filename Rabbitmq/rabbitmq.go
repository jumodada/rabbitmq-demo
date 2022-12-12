package main

import "github.com/streadway/amqp"

const MQUrl = "amqp://admin:123@127.0.0.1:5672/test"

type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	QueName  string
	Exchange string
	Key      string
	MqUrl    string
}

func newRabbit(queueName string, exchange string) *RabbitMQ {
	return &RabbitMQ{QueName: queueName, Exchange: exchange}
}

func main() {

}
