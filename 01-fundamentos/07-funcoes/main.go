package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(1, 2))

	valor, err := sum3(1, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}

	valor, err = sum3(51, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}

}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (int, bool) {
	if a+b > 50 {
		return a + b, true
	}

	return a + b, false
}

func sum3(a int, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
