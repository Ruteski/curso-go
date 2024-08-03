package main

import (
	"eventos/pkg/rabbitmq"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "minhaFila")
	for msg := range msgs {
		fmt.Printf("Received a message: %s\n", msg.Body)
		msg.Ack(false) // mensagem ja foi lida e nao é pra colocar ela na fila novamente
	}
}
