package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// c := http.Client{Timeout: time.Second}
	// c := http.Client{Timeout: 5 * time.Second}
	// c := http.Client{Timeout: time.Duration(1) * time.Second}
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}
