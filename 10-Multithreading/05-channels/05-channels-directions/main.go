package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

// thread 1 - o main sempre é a primeira thread do sistema
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
