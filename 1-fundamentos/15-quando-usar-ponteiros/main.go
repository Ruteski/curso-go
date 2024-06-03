﻿package main

func soma(a, b int) int {
	a = 50
	return a + b
}

func soma2(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println(soma(minhaVar1, minhaVar2))
	println(minhaVar1)

	println(soma2(&minhaVar1, &minhaVar2))
	println(minhaVar1)
}
