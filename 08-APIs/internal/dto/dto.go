package dto

// type  CreateProductDTO struct {}

type CreateProductInput struct {
	Nome  string  `json:"nome"`
	Price float32 `json:"price"`
}
