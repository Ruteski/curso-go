package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type message struct {
	id  int64
	msg string
}

func main() {
	//c1 := make(chan int)
	//c2 := make(chan int)
	c1 := make(chan message)
	c2 := make(chan message)
	var i int64 = 0

	// lendo mensagens do RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1) // para nao ter problema de concorrencia
			time.Sleep(3 * time.Second)
			msg := message{i, "hello from rabbitmq"}
			c1 <- msg
		}

	}()

	//lendo mensagens do Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1) // para nao ter problema de concorrencia
			time.Sleep(4 * time.Second)
			msg := message{i, "hello from kafka"}
			c2 <- msg
		}

	}()

	//for i := 0; i < 3; i++ {
	//colocando assim, ele fica esperando receber mensagem para sempre
	//for {
	//	select {
	//	case msg1 := <-c1: //imagine recebendo msg do rabbitmq
	//		println("receive c1:", msg1)
	//	case msg2 := <-c2: //imagine recebendo msg do kafka
	//		println("receive c2:", msg2)
	//	case <-time.After(5 * time.Second):
	//		println("timeout")
	//		//default:
	//		//	println("caso nenhum acima funcione, eu executo")
	//	}
	//}

	for {
		select {
		case msg := <-c1: //imagine recebendo msg do rabbitmq
			fmt.Printf("Received from RabbitMQ ID: %d - message: %s\n", msg.id, msg.msg)
		case msg := <-c2: //imagine recebendo msg do kafka
			fmt.Printf("Received from Kafka ID: %d - message: %s\n", msg.id, msg.msg)
		case <-time.After(5 * time.Second):
			println("timeout")
			//default:
			//	println("caso nenhum acima funcione, eu executo")
		}
	}

	//select {
	//case msg1 := <-c1:
	//	println("receive c1:", msg1)
	//case msg2 := <-c2:
	//	println("receive c2:", msg2)
	//case <-time.After(5 * time.Second):
	//	println("timeout")
	//	//default:
	//	//	println("caso nenhum acima funcione, eu executo")
	//}

}
