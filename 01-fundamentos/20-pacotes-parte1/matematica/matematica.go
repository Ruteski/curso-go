package matematica

import "fmt"

func Soma[T int | float64](a, b T) T {
	return a + b
}

func soma2[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
	cor   string
}

func (carro Carro) Andar() {
	fmt.Println("Carro andando")
}

type carro struct {
	Marca string
}
