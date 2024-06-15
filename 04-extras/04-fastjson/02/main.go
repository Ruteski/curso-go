// olhar a doc a funcao scanner, que pega json ferrado e tenta ler os dados

package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{"user": {"name": "Jhon Doe", "age": 30}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := value.GetObject("user")
	fmt.Printf("user name: %s\n", user.Get("name"))
	fmt.Printf("user age: %s\n", user.Get("age")) // pego como string, pq msm sendo int, no json ele ainda é uma string

	fmt.Println("============================================")

	// com struct e unmarshal
	userJson := value.Get("user").String()
	var user2 User
	if err = json.Unmarshal([]byte(userJson), &user2); err != nil {
		panic(err)
	}
	fmt.Printf("user name: %s\n", user2.Name)
	fmt.Printf("user age: %d\n", user2.Age)
}
