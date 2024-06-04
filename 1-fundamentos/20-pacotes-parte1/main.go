package main

// go mod init <nome> | colocar o link do github, ajuda as pessoas baixarem os modulos que fazemos
// go mod tidy - atualiza os pacotes

import (
	"fmt"
	"github.com/ruteski/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v \n", s)

	println(matematica.A)
	
	carro := matematica.Carro{Marca: "audi"}
	carro.Andar()
}
