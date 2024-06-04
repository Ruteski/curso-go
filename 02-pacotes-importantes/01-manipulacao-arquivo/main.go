package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("hello, world!") // grava string
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo")) // grava realmente bytes | []byte -> slice de bytes
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	///// Leitura de arquivo
	//arquivo, err := os.Open("arquivo.txt")
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de arquivo pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10) //o buffer vai ser um slice de 10 em 10 bytes
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// removendo arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
