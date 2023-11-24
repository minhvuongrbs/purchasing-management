package product

import (
	"context"
	"errors"
)

var ErrRepositoryGetProductsNotFound = errors.New("repository get products not found")

type Repository interface {
	CreateProduct(ctx context.Context, product Product) error
	GetProducts(ctx context.Context) ([]Product, error)
}
