package main

func main() {
	// Memoria —> Endrego —> Valor
	// variavel -> ponteiro que tem um endereco na memora —> valor
	a := 10
	println(a)
	println(&a) // printa o endereco de memoria de 'a'

	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	println(*ponteiro)
	println("Novo valor de 'a': ", a)

	b := &a
	println(b)
	println(*b) // *b chamado de direfence
	*b = 30

	println("valor de a:", a)
	println("valor de a na memoria:", &a)
	println("valor de b:", b)

}
