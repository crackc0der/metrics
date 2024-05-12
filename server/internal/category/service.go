package category

import (
	"context"
	"fmt"

	product_grpc "server/api/note_v1"
)

type RepositoryInterface interface {
	SelectCategories(_ context.Context) (*product_grpc.AllCategoryMessage, error)
	InsertCategory(_ context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error)
	UpdateCategory(_ context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error)
	DeleteCategory(_ context.Context, id *product_grpc.CategoryRequest) (*product_grpc.CategoryResponse, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetCategories(ctx context.Context) (*product_grpc.AllCategoryMessage, error) {
	categories, err := s.repository.SelectCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in service's method GetCategories: %w", err)
	}

	return categories, nil
}

func (s *Service) AddCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := s.repository.InsertCategory(ctx, cat)
	if err != nil {
		return nil, fmt.Errorf("error in service's method AddCategory: %w", err)
	}

	return category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := s.repository.UpdateCategory(ctx, cat)
	if err != nil {
		return nil, fmt.Errorf("error in service's method UpdateCategory: %w", err)
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id *product_grpc.CategoryRequest) (*product_grpc.CategoryResponse, error) {
	result, err := s.repository.DeleteCategory(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in service's method DeleteCategory: %w", err)
	}

	return result, nil
}
