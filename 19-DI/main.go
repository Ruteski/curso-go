// pra rodar com o wire tem que ser esse comando: go run main.go wire_gen.go

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//repository := product.NewProductRepository(db)
	//usecase := product.NewProductUseCase(repository)
	usecase := NewUseCase(db)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)
}
