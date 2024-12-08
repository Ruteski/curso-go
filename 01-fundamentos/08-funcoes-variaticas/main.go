package main

import "fmt"

func main() {
	fmt.Println("Primeiro sum: ", sum(1, 3))
	fmt.Println("Segundo sum: ", sum(1, 3, 2, 6, 8, 9, 34, 5, 335, 78, 6543, 4, 6, 76, 67, 2))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
