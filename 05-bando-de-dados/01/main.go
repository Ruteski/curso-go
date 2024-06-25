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
	product.ID = "751956d8-3e6d-40ba-9f2f-525b22113897"
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("Produto alterado com sucesso!")

	fmt.Println("=======================================")

	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Producto: %+v possui o valor de R$%.2f\n", p.Name, p.Price)

	fmt.Println("=======================================\n")

	products, err := selectAllProducts(db)

	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Product: %v, possui o valor de %.2f\n", p.Name, p.Price)
	}

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

func selectProduct(db *sql.DB, id string) (*Product, error) {
	//func selectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	// err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price) // usando context para determinar um tempo para retornar a consulta
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
