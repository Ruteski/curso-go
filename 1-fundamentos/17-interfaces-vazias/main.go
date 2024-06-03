package main

import "fmt"

// parecido com tipo generico <T>
func main() {
	var x interface{} = 10
	var y interface{} = "hello world"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v \n", t, t)
}
