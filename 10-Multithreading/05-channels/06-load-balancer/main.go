package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(1 * time.Second)
	}
}

// thread 1 - o main sempre é a primeira thread do sistema
func main() {
	data := make(chan int)
	qtdWorkers := 10000 //10000 // 90

	// inicializa os workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100000; i++ {
		data <- i
	}
}
