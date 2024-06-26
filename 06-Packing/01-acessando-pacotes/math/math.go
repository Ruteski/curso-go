package math

var X string = "hello world"

type Math struct {
	a int
	b int
}

type math struct {
	a int
	b int
}

func NewMath(a, b int) Math {
	return Math{a: a, b: b}
}

func Newmath(a, b int) math {
	return math{a: a, b: b}
}

func (m *Math) Add() int {
	return m.a + m.b
}
