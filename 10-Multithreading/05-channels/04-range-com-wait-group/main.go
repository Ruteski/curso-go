package main

import (
	"fmt"
	"sync"
)

// thread 1 - o main sempre é a primeira thread do sistema
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10) // eu sei que a iteracao é de 10
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// agora nao preciso, pq estou usando o waitgroup
	//close(ch) // indica que nada mais vai entrar no canal, para nao ocorrer deadlock
}
