package main

import (
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileserver)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow Blog"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
