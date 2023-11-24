package product

import (
	"context"
	"database/sql"
	"purchasing-management/internal/adapters/repository/da"
	"purchasing-management/internal/domain/product"

	"github.com/pkg/errors"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) CreateProduct(ctx context.Context, p product.Product) error {
	q := da.New(r.db)
	affectedRows, err := q.CreateProduct(ctx, p.Name)
	if err != nil {
		return errors.Wrap(err, "create product error")
	}
	if affectedRows == 0 {
		return errors.New("empty affected rows")
	}
	return nil
}

func (r Repository) GetProducts(ctx context.Context) ([]product.Product, error) {
	q := da.New(r.db)
	rows, err := q.GetProducts(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return []product.Product{}, product.ErrRepositoryGetProductsNotFound
	}
	products := make([]product.Product, 0)
	for _, row := range rows {
		products = append(products, product.Product{ID: row.ID, Name: row.Name})
	}

	return products, nil
}
