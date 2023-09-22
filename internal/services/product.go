package services

import (
	"context"
	"savannah/internal/models"
	"savannah/internal/repository/db"
)

type ProductService interface {
	ProductCreate(ctx context.Context, product models.Product) error
	ProductUpdate(ctx context.Context, product models.Product) error
	ProductDelete(ctx context.Context, product models.Product) error
	ProductGet(ctx context.Context, productID int) (*models.Product, error)
	ProductsList(ctx context.Context) ([]*models.Product, error)
}

func (s *SQLStore) ProductCreate(ctx context.Context, product models.Product) error {
	if err := s.store.CreateProduct(ctx, db.CreateProductParams{
		Name:  product.Name,
		Price: int32(product.Price),
	}); err != nil {
		return err
	}
	return nil
}
func (s *SQLStore) ProductUpdate(ctx context.Context, product models.Product) error {
	if err := s.store.UpdateProduct(ctx, db.UpdateProductParams{
		ID:    int32(product.ID),
		Name:  product.Name,
		Price: int32(product.Price),
	}); err != nil {
		return err
	}
	return nil
}
func (s *SQLStore) ProductDelete(ctx context.Context, product models.Product) error {
	if err := s.store.DeleteProduct(ctx, int32(product.ID)); err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) ProductGet(ctx context.Context, productID int) (*models.Product, error) {
	product, err := s.store.GetProduct(ctx, int32(productID))
	if err != nil {
		return nil, err
	}
	return &models.Product{
		ID:        int64(product.ID),
		CreatedAt: product.CreatedAt.Time,
		UpdatedAt: product.UpdatedAt.Time,
		DeletedAt: product.DeletedAt.Time,
		Name:      product.Name,
		Price:     int(product.Price),
	}, nil
}

func (s *SQLStore) ProductsList(ctx context.Context) ([]*models.Product, error) {
	products, err := s.store.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	productsList := make([]*models.Product, len(products))
	for i, v := range products {
		productsList[i] = &models.Product{
			ID:        int64(v.ID),
			CreatedAt: v.CreatedAt.Time,
			UpdatedAt: v.UpdatedAt.Time,
			DeletedAt: v.DeletedAt.Time,
			Name:      v.Name,
			Price:     int(v.Price)}
	}

	return productsList, nil
}
