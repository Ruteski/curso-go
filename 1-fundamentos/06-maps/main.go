package main

import "fmt"

func main() {
	//key(string) - value(int)
	salarios := map[string]int{"Lincoln": 110520, "Noah": 78500, "Melanie": 12500}
	delete(salarios, "Melanie")

	salarios["Mel"] = 5000
	fmt.Println(salarios)

	salarios["Melanie"] = 7850
	fmt.Println(salarios)

	fmt.Println(salarios["Lincoln"])

	sal := make(map[string]int)
	sal["Lincoln"] = 110520
	fmt.Println(sal)

	for nome, salario := range salarios {
		fmt.Printf("o salario de %s é de %d \n", nome, salario)
	}

	for _, salario := range salarios { // o underline se chama blankidentifier
		fmt.Printf("o salario é de %d \n", salario)
	}

}
