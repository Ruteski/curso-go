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

type Cliente1 struct {
	Nome    string
	Idade   int
	Ativo   bool
	address Endereco
}

func main() {
	lincoln := Cliente{
		Nome:  "Lincoln",
		Idade: 40,
		Ativo: true,
	}

	lincoln.Cidade = "Curitiba"
	lincoln.Endereco.Estado = "Parana"
	fmt.Println(lincoln)

	lincoln.Endereco = Endereco{
		Rua:    "antonio",
		Numero: 10,
		Cidade: "Curitiba",
		Estado: "Parana",
	}
	fmt.Println("\n", lincoln)

	lincoln2 := Cliente1{
		Nome:  "Noah",
		Idade: 5,
		Ativo: false,
	}

	lincoln2.address.Cidade = "Curitiba"
	lincoln2.address.Estado = "Parana"
	lincoln2.address.Rua = "vidolin"
	lincoln2.address.Numero = 125
	fmt.Println("\n", lincoln2)
}
