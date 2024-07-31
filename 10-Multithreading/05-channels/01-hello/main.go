package main

import "fmt"

// thread 1 - o main sempre é a primeira thread do sistema
func main() {
	// canal vazio
	canal := make(chan string) // criando um canal de strings

	// thread 2
	go func() {
		// canal cheio
		canal <- "Olá mundo!"
	}()

	// na thread 1 novamente
	// canal esvaziou
	// peguei o valor que estava na thread 2 e passei ele para a thread 1 - comunicacao entre threads
	msg := <-canal
	fmt.Println(msg)
}
