package product

import (
	"context"
	"fmt"

	product_grpc "server/api/note_v1"
)

type RepositoryInterface interface {
	SelectProducts(_ context.Context) (*product_grpc.AllProductMessage, error)
	SelectProductByID(_ context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error)
	InsertProduct(_ context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
	DeleteProductByID(_ context.Context, productID *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error)
	UpdateProduct(_ context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetProducts(ctx context.Context) (*product_grpc.AllProductMessage, error) {
	products, err := s.repository.SelectProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting products in service's method GetProducts: %w", err)
	}

	return products, nil
}

func (s *Service) GetProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	product, err := s.repository.SelectProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in service.GetProduct: %w", err)
	}

	return product, nil
}

func (s *Service) AddProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := s.repository.InsertProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in service.AddProduct: %w", err)
	}

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, productID *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	res, err := s.repository.DeleteProductByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("error in service.DeleteProduct: %w", err)
	}

	return &product_grpc.ProductResponse{Deleted: res.GetDeleted()}, nil
}

func (s *Service) UpdateProduct(ctx context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := s.repository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error in service.UpdateProduct: %w", err)
	}

	return product, nil
}
