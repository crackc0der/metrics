package category

import (
	product_grpc "client/api/note_v1"
	"context"
	"fmt"
)

type Repository struct {
	client product_grpc.CategoryClient
}

func NewRepository(cl product_grpc.CategoryClient) *Repository {
	return &Repository{client: cl}
}

func (r *Repository) SelectCategories(ctx context.Context) (*product_grpc.AllCategoryMessage, error) {
	categoris, err := r.client.GetCategories(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method SelectCategories: %w", err)
	}

	return categoris, nil
}

func (r *Repository) InsertCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := r.client.AddCategory(ctx, cat)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method InsertCategory: %w", err)
	}

	return category, nil
}

func (r *Repository) UpdateCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := r.client.UpdateCategory(ctx, cat)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method UpdateCategory: %w", err)
	}

	return category, nil
}

func (r *Repository) DeleteCategory(ctx context.Context, id *product_grpc.CategoryRequest) (*product_grpc.CategoryResponse, error) {
	result, err := r.client.DeleteCategory(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method DeleteCategory: %w", err)
	}

	return result, nil
}
