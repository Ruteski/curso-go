package main

import "fmt"

type Cliente struct {
	Nome string
}

func (cliente Cliente) andou() {
	cliente.Nome = "Lincoln Ruteski"
	fmt.Printf("O cliente %s andou \n", cliente.Nome)
}

type Conta struct {
	saldo int
}

func (conta *Conta) simular(valor int) int {
	conta.saldo += valor
	println("Valor dentro do simular: ", conta.saldo)
	return conta.saldo
}

func NewConta() *Conta {
	return &Conta{0}
}

func main() {
	lincoln := Cliente{
		Nome: "Lincoln",
	}

	lincoln.andou()

	fmt.Printf("O valor da struct com nome %v \n\n\n", lincoln.Nome)

	conta := Conta{saldo: 105}
	println(conta.simular(250))
	println("Valor do meu saldo, depois de SIMULAR: ", conta.saldo)
}
