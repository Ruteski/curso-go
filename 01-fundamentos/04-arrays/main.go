package main

import (
	"fmt"
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// busca o ultimo valor do array
	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])

	// for indice, valor
	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é %d \n", i, v)
	}
}
