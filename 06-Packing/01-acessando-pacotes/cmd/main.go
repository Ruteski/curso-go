package main

import (
	"fmt"
	"github.com/Ruteski/curso-go/math"
)

func main() {
	// m := math.NewMath(1, 2)
	//=== ou assim \/
	m := math.Math{}
	//m2 := math.Math{}
	m2 := math.NewMath(1, 2)

	fmt.Println(m.Add())
	fmt.Println(math.X)

	fmt.Println(m2.Add())
	fmt.Println(math.X)
}
