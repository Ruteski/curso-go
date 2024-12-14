package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{repository: repository}
}

// GetProduct apenas para exemplo, pois aqui deveria retornar um dto e nao a entidade
func (p *ProductUseCase) GetProduct(id int) (Product, error) {
	return p.repository.GetProduct(id)
}
