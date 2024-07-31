// baixar o apache benchmark - ab - para fazer o teste

package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

// aqui é simples concorrencia
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number += 1
		// w.Write([]byte("Você teve acesso a página " + strconv.Itoa(int(number)) + " vezes"))
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})

	http.ListenAndServe(":3000", nil)
}
