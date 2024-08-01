/*
	Eu fiquei na dúvida do que era pra ser feito no desafio, abri tópico no fórum, mas não fui respondido,
	não entendi se era pra fazer apenas a parte de multithreading com as especificações do desafio ou se
   era pra fazer uma api com multithreading, o enunciado me deixou na duvida, entao fiz com a api, mas
   caso não fosse pra ser api, também tenho o codigo apenas da parte de multithreading, se necessário
   eu subo esse codigo também.
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
	http.HandleFunc("/", buscaCEP)
	fmt.Println("Server online na porta 8000 🚀")
	http.ListenAndServe(":8000", nil)
}

func buscaCEP(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c1 := make(chan response)
	c2 := make(chan response)

	go func() {
		//time.Sleep(1 * time.Second)
		req, err := http.Get("http://viacep.com.br/ws/" + cepParam + "/json/")
		if err != nil {
			log.Fatal(err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		c2 <- response{"https://viacep.com.br", string(res)}
	}()

	go func() {
		//time.Sleep(2 * time.Second)
		req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cepParam)
		if err != nil {
			log.Fatal(err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		c1 <- response{"https://brasilapi.com.br", string(res)}
	}()

	select {
	case msg := <-c1:
		fmt.Printf("Site de consulta: %s.\nDados do endereço: %s\n\n", msg.site, msg.res)
	case msg := <-c2:
		fmt.Printf("Site de consulta: %s.\nDados do endereço: %s\n\n", msg.site, msg.res)
	case <-time.After(time.Second * 1):
		println("timeout")
	}
}
