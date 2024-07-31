package main

import (
	"fmt"
)

// thread 1 - o main sempre é a primeira thread do sistema
func main() {
	ch := make(chan int)
	go publish(ch)
	//reader(ch) // aqui nao tem goroutine para o programa nao morrer na primeira passada

	//mesma coisa que fazer na funcao reader
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch) // indica que nada mais vai entrar no canal, para nao ocorrer deadlock
}
