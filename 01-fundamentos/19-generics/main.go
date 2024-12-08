package main

type MyNumber int

// constraint - crio um tipo
type Number interface {
	// o ~ na frente do int é pra poder usar o tipo MyNumber, se nao tiver isso ele grita, pq int != MyNumber
	// exemplo o m3 usando ao SomaConstraint
	~int | float64
}

func SomaConstraint[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

// func Compara2[T any](a T, b T) bool {
func Compara2[T comparable](a T, b T) bool {
	// o comparable só aceita igualdade "=="
	if a == b {
		return true
	}
	return false
}

// ver doc - https://pkg.go.dev/golang.org/x/exp/constraints
// possui varias constraints prontas

func main() {
	m := map[string]int{"lincoln": 1000, "Noah": 2000, "Melanie": 3000}
	m2 := map[string]float64{"lincoln": 1000.0, "Noah": 2000.0, "Melanie": 3000.0}
	println(Soma(m))
	println(Soma(m2))

	//usando a func de constraint
	println(SomaConstraint(m))
	println(SomaConstraint(m2))

	//===============
	m3 := map[string]MyNumber{"lincoln": 1000, "Noah": 2000, "Melanie": 3000}
	println(SomaConstraint(m3))

	//===============
	println(Compara(10, 10.0))
	println(Compara2(10, 10.0))
}
