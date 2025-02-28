﻿/*
	Eu fiquei na dúvida do que era pra ser feito no desafio, abri tópico no fórum, mas a resposta continuou me deixando,
	na duvida não entendi se era pra fazer apenas a parte de multithreading com as especificações do desafio ou se
   era pra fazer uma api com multithreading, o e nunciado me deixou na duvida, entao fiz com a api também, subi aqui no
	github
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type response struct {
	site string
	res  string
}

func main() {
	channel1 := make(chan response)
	channel2 := make(chan response)
	cep := "83010100"

	go func() {
		//time.Sleep(1 * time.Second)
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			log.Fatal(err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		channel1 <- response{"https://viacep.com.br", string(res)}
	}()

	go func() {
		//time.Sleep(2 * time.Second)
		req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
		if err != nil {
			log.Fatal(err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		channel2 <- response{"https://brasilapi.com.br", string(res)}
	}()

	select {
	case msg := <-channel1:
		fmt.Printf("Site de consulta: %s.\nDados do endereço: %s\n\n", msg.site, msg.res)
	case msg := <-channel2:
		fmt.Printf("Site de consulta: %s.\nDados do endereço: %s\n\n", msg.site, msg.res)
	case <-time.After(time.Second * 3):
		println("timeout")
	}
}
