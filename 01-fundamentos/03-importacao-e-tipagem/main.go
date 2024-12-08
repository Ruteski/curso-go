package main

import "fmt"

type ID int

var e = 1.2

func main() {
	var a ID = 20

	// %T - traz o tipo da variavel
	fmt.Printf("O tipo de 'E' é %T", e)
	fmt.Println("")

	// %v - traz o valor da variavel
	fmt.Printf("O valor de 'E' é %v", e)
	fmt.Println("")
	fmt.Println("")

	fmt.Printf("O tipo de 'A' é %T", a)
	fmt.Println("")

	fmt.Printf("O valor de 'A' é %v", a)
	fmt.Println("")
}
