package category

import (
	"context"
	"time"

	"github.com/Furkan-Gulsen/Checkout-System/internal/entities"
)

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (p *CategoryService) List() ([]entities.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return p.categoryRepository.List(ctx)
}

func (p *CategoryService) Create(category entities.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return p.categoryRepository.Create(ctx, category)
}
