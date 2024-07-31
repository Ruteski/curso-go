// baixar o apache benchmark - ab - para fazer o teste

// usado pra verificar se o meu programa esta tendo problema de concorrencia(race condition)
// é mais lento e dev ser usado apenas em dev para verificar o sistema
// go run -race main.go

package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

// aqui é simples concorrencia
func main() {
	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		//number += 1
		atomic.AddUint64(&number, 1) // pra nao precisar ficar colocando lock e unlock o tempo todo
		// m.Unlock()

		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})

	http.ListenAndServe(":3000", nil)
}
