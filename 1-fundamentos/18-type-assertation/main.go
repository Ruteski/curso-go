package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Lincoln Ruteski"
	println(minhaVar)
	println(minhaVar.(string)) // isso é o type assertation, confirmar que a variavel dé de um certo tipo

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v \n", res, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)

}
