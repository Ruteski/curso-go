package main

import "fmt"

func main() {
	x := 10
	for i := range x {
		fmt.Println(i)
	}

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, v := range values {
		// v := v // assim que arrumavam antes da versao 1.22
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}
