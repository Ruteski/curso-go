package main

// GAMBIARRA - rs
// quando no pacote nao estiver publicado e só estiver na maquina, pode ser usado dessa forma, mas nao mto recomendado, tem o workspaces
// 1 - go mod edit -replace curso-go/math=../math
// 2 - go mod tidy

import "curso-go/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
