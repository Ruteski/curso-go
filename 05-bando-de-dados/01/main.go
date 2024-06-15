package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // _ blank identifier
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Notebook", 1899.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Produto inserido com sucesso!")

	product.Price = 7570.00
	product.ID = "567037a6-5a8a-436c-b423-15d07c5b739c"
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Produto alterado com sucesso!")
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price) // ao inves do "_, err" da pra usar "res, err"(res = result)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}
