package main

import (
	"fmt"
)

// funcoes anonimas
func main() {
	// comumente usado para executar funcoes sem retorno
	func() {
		//	rodar_webserver()
	}()

	fmt.Println(sum(1, 3, 2, 6, 8, 9, 34, 5, 335, 78, 6543, 4, 6, 76, 67, 2))

	total := func() int {
		return sum(1, 3, 2, 6, 8, 9, 34, 5, 335, 78, 6543, 4, 6, 76, 67, 2) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
