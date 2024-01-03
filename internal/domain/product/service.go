package product

import "github.com/Furkan-Gulsen/Checkout-System/internal/entities"

type ProductService struct {
	ProductRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (p *ProductService) List() ([]entities.Item, error) {
	return p.ProductRepository.List()
}

func (p *ProductService) Create(product entities.Item) error {
	return p.ProductRepository.Create(product)
}

func (p *ProductService) Update(product entities.Item) error {
	return p.ProductRepository.Update(product)
}

func (p *ProductService) Delete(product entities.Item) error {
	return p.ProductRepository.Delete(product)
}
