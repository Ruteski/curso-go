//go:build wireinject
// +build wireinject

package main

import (
	"19-DI/product"
	"database/sql"
	"github.com/google/wire"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
