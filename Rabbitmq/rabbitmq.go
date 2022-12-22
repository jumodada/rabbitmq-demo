package Rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const MQUrl = "amqp://admin:123@127.0.0.1:5672/test"

type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	QueName  string
	Exchange string
	Key      string
	MqUrl    string
}

func NewRabbit(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueName: queueName, Exchange: exchange, Key: key, MqUrl: MQUrl}
}

func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func rabbitSimpleMode(queueName string) *RabbitMQ {
	rabbitmq := NewRabbit(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnError(err, "connected error")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "getting channel error")
	return rabbitmq
}

func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.channel.QueueDeclare(r.QueName,
		// 是否持久化
		false,
		//是否自动删除
		false,
		// 是否具有排他性
		false,
		//是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	r.channel.Publish(
		r.Exchange,
		r.QueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

}

func (r *RabbitMQ) ConsumeSimple() {
	_, err := r.channel.QueueDeclare(r.QueName,
		// 是否持久化
		false,
		//是否自动删除
		false,
		// 是否具有排他性
		false,
		//是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	messages, err := r.channel.Consume(
		r.QueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a messsgae: %s", d.Body)
			fmt.Println(d.Body)
		}
	}()

	log.Printf("[*] waiting for messages, To exit press Ctrl + C")
	<-forever
}
