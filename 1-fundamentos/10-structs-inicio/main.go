package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	lincoln := Cliente{
		Nome:  "Lincoln",
		Idade: 40,
		Ativo: true,
	}
	fmt.Printf("Nome: %s \n Idade: %d\n Ativo: %t\n", lincoln.Nome, lincoln.Idade, lincoln.Ativo)

	lincoln.Ativo = false
	fmt.Printf("Nome: %s \n Idade: %d\n Ativo: %t\n", lincoln.Nome, lincoln.Idade, lincoln.Ativo)
}
