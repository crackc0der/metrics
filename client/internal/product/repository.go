package product

import (
	product_grpc "client/api/note_v1"
	"context"
	"fmt"
)

type Repository struct {
	client product_grpc.ProductClient
}

func NewRepository(cl product_grpc.ProductClient) *Repository {
	return &Repository{client: cl}
}

func (r *Repository) SelectProducts(ctx context.Context) (*product_grpc.AllProductMessage, error) {
	products, err := r.client.GetProducts(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error in repository.SelectProducts: %w", err)
	}

	return products, nil
}

func (r *Repository) SelectProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	product, err := r.client.GetProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in repository.SelectProduct: %w", err)
	}

	return product, nil
}

func (r *Repository) InsertProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := r.client.AddProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in repository.InsertProduct: %w", err)
	}

	return product, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	result, err := r.client.DeleteProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in repository.DeleteProduct: %w", err)
	}

	return result, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := r.client.UpdateProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in repository.UpdateProduct: %w", err)
	}

	return product, nil
}
