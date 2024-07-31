// Sempre que eu for trabalhar com goroutines, sempre tera que ter um processo segurando o sistema ate a finalização
// de todas as goroutines, pois se o sistema sair antes, tudo para
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done() // informa ao waitgroup que finalizou o processo
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}

	// no final tenho 25 operacoes, 10 de A, 10 de B e 5 anonymous
	waitGroup.Add(25)

	// Thread 2
	go task("A", &waitGroup)

	// Thread 3
	go task("B", &waitGroup)

	// Thread 4 - funcao anonima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: task %s is running\n", i, "anonynous")
			time.Sleep(1 * time.Second)
			waitGroup.Done() // informa ao waitgroup que finalizou o processo
		}
	}()
	waitGroup.Wait() // espera ate as 25 threads terminarem
}
