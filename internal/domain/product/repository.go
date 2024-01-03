package product

import "github.com/Furkan-Gulsen/Checkout-System/internal/entities"

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) List() ([]entities.Item, error) {
	return nil, nil
}

func (p *ProductRepository) Create(product entities.Item) error {
	return nil
}

func (p *ProductRepository) Update(product entities.Item) error {
	return nil
}

func (p *ProductRepository) Delete(product entities.Item) error {
	return nil
}
