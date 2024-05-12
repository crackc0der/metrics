package product

import (
	product_grpc "client/api/note_v1"
	"context"
	"fmt"
)

type RepositoryInterface interface {
	SelectProducts(ctx context.Context) (*product_grpc.AllProductMessage, error)
	SelectProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error)
	InsertProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
	DeleteProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error)
	UpdateProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
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
		return nil, fmt.Errorf("error in service.GetProducts: %w", err)
	}

	return products, nil
}

func (s *Service) GetProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	product, err := s.repository.SelectProduct(ctx, id)
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

func (s *Service) DeleteProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	result, err := s.repository.DeleteProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in service.DeleteProduct: %w", err)
	}

	return result, nil
}

func (s *Service) UpdateProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := s.repository.UpdateProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in service.UpdateProduct: %w", err)
	}

	return product, nil
}
