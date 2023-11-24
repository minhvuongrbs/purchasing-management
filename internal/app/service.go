package app

import (
	"context"
	"errors"
	"purchasing-management/internal/domain/product"
)

type OrderService struct {
	repo product.Repository
}

func NewOrderService(repository product.Repository) OrderService {
	return OrderService{repo: repository}
}

func (s OrderService) GetProducts(ctx context.Context) ([]product.Product, error) {
	products, err := s.repo.GetProducts(ctx)
	if errors.Is(err, product.ErrRepositoryGetProductsNotFound) {
		return []product.Product{}, nil
	}
	if err != nil {
		return []product.Product{}, err
	}

	return products, nil
}

func (s OrderService) CreateProduct(ctx context.Context, request CreateProductRequest) error {
	err := s.repo.CreateProduct(ctx, product.Product{Name: request.Name})
	return err
}
