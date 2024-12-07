package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
}

func main() {
	lincoln := Cliente{
		Nome:  "Lincoln",
		Idade: 40,
		Ativo: true,
		Endereco: Endereco{
			Rua:    "antonio",
			Numero: 10,
			Cidade: "Curitiba",
			Estado: "Parana",
		},
	}

	fmt.Println(lincoln)
	lincoln.Desativar()
	fmt.Println(lincoln)

}
