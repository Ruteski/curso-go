package main

import "fmt"

// a interface só me permite passar metodos
type Pessoa interface {
	Desativar()
}

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

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
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

	minhaEmpresa := Empresa{}

	Desativacao(lincoln)
	Desativacao(minhaEmpresa)
}
