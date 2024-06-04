package main

import "fmt"

func main() {
	//req, err := http.Get("http://www.google.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer req.Body.Close()
	//
	//res, err := io.ReadAll(req.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(res))

	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
