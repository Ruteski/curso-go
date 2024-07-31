package main

// thread 1 - o main sempre é a primeira thread do sistema
func main() {
	//canal vazio
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		// fazendo isso, nao ocorre o deadlock no final
		forever <- true
	}()

	// se fizer assim sem ter a goroutine acima, vai dar deadlock tb
	//forever <- true

	// canal esperando estar cheio para esvaziar
	<-forever
	// da forma que esta acima, sem nada ou com o um goroutine, vai ocorrer deadlock,
	// porem com a goroutine no meio, o <-forever segura o processamento do sistema, ate todas as threads terminarem
}
