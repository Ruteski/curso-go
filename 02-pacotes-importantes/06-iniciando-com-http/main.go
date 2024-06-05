package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCEP)
	http.HandleFunc("/new", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Hello World 2, com func anonima"))
	})
	http.ListenAndServe(":8080", nil) //isso cria um http server
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
