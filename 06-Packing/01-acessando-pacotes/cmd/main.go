package main

import (
	"fmt"
	"github.com/Ruteski/curso-go/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
