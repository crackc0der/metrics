package product

import (
	"context"
	"fmt"

	// Importing pgx for indirect use via another package.
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"

	product_grpc "server/api/note_v1"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectProducts(_ context.Context) (*product_grpc.AllProductMessage, error) {
	query := "SELECT * FROM product"

	var products []*product_grpc.ProductMessage

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("error in repository.SelectProducts: %w", err)
	}

	allProducts := &product_grpc.AllProductMessage{
		Products: products,
	}

	return allProducts, nil
}

func (r *Repository) SelectProductByID(_ context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	query := "SELECT * FROM product WHERE product_id=$1"

	var product product_grpc.ProductMessage

	err := r.db.Get(&product, query, id.GetId())
	if err != nil {
		return nil, fmt.Errorf("error repository.SelectProductById: %w", err)
	}

	return &product, nil
}

func (r *Repository) InsertProduct(_ context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	query := `INSERT INTO product (product_name, product_category_id, product_price) VALUES($1, $2, $3) 
		RETURNING product_id, product_name, product_category_id, product_price`

	var product product_grpc.ProductMessage

	err := r.db.QueryRowx(query, prod.GetProductName(), prod.GetProductCategoryID(), prod.GetProductPrice()).StructScan(&product)
	if err != nil {
		return nil, fmt.Errorf("error repository.InsertProduct: %w", err)
	}

	return &product, nil
}

func (r *Repository) DeleteProductByID(_ context.Context, productID *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	query := "DELETE FROM product WHERE product_id=$1"

	_, err := r.db.Exec(query, productID.GetId())
	if err != nil {
		return nil, fmt.Errorf("error repository.DeleteProductById: %w", err)
	}

	return &product_grpc.ProductResponse{Deleted: true}, nil
}

func (r *Repository) UpdateProduct(_ context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	query := `UPDATE product SET product_name=$1, product_category_id=$2, product_price=$3 WHERE product_id=$4
			RETURNING product_id, product_name, product_category_id, product_price`

	var updatedProduct product_grpc.ProductMessage

	err := r.db.QueryRowx(query, product.GetProductName(), product.GetProductCategoryID(), product.GetProductPrice(), product.GetId()).StructScan(&updatedProduct)
	if err != nil {
		return nil, fmt.Errorf("error repository.UpdateProduct: %w", err)
	}

	return &updatedProduct, nil
}
