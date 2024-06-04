package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`                // com tag json n, posso ignorar o campo usando json:"-"
	Saldo  int `json:"saldo" validate:"gt=0"` // com tag json s
}

func main() {
	conta := Conta{Numero: 100, Saldo: 200}
	res, err := json.MarshalIndent(conta, "", "  ") // transforma em json, sempre retorna em []bytes
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(res))

	// usando encoder
	// json.NewEncoder(os.Stdout).Encode(conta)  // apenas exibi o json
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Println("error:", err)
	}

	//=============== transformando json em struct
	jsonPuro := []byte(`{"saldo":200, "numero":2}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Saldo: ", contaX.Saldo)   //depois de usar tag, o retorno vai ser 0
	fmt.Println("Numero: ", contaX.Numero) //depois de usar tag, o retorno vai ser 0

	//=========== usando tag pra mapear
	jsonPuro2 := []byte(`{"s":200, "n":2}`)
	var contaX2 Conta
	err = json.Unmarshal(jsonPuro2, &contaX2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Saldo: ", contaX2.Saldo)

}
