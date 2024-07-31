// Sempre que eu for trabalhar com goroutines, sempre tera que ter um processo segurando o sistema ate a finalização
// de todas as goroutines, pois se o sistema sair antes, tudo para
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")

	// Thread 3
	go task("B")

	// Thread 4
	go task("C")

	// Thread 5
	go task("D")

	// Thread 6
	go task("E")

	// Thread 7
	go task("F")

	// Thread 8
	go task("G")

	// Thread 9 - funcao anonima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: task %s is running\n", i, "anonynous")
			time.Sleep(1 * time.Second)
		}
	}()

	// aqui nao tem nada depois da thread 3, go chega aqui e sai.

	// pra teste
	time.Sleep(15 * time.Second)
}
