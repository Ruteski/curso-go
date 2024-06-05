package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World"))
	//})
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blogStruct{})
	mux.Handle("/blog2", blogStruct2{title: "My Blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Ruteski"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blogStruct struct{}

func (blog blogStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog"))
}

type blogStruct2 struct {
	title string
}

func (blog blogStruct2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(blog.title))
}
