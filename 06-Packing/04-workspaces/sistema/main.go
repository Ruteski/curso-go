package main

import (
	"curso-go/math"
	"github.com/google/uuid"
)

// go mod tidy -e

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())

	println(uuid.New().String())
}
