package product

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"

	product_grpc "server/api/note_v1"
)

type ServiceInterface interface {
	GetProducts(ctx context.Context) (*product_grpc.AllProductMessage, error)
	GetProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error)
	AddProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
	DeleteProduct(ctx context.Context, productID *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error)
	UpdateProduct(ctx context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
	product_grpc.UnimplementedProductServer
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	//nolint
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetProducts(ctx context.Context, _ *emptypb.Empty) (*product_grpc.AllProductMessage, error) {
	products, err := e.service.GetProducts(ctx)
	if err != nil {
		e.log.Error("error in endpoint.GetProducts: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.GetProducts: %w", err)
	}

	return products, nil
}

func (e *Endpoint) GetProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	product, err := e.service.GetProduct(ctx, id)
	if err != nil {
		e.log.Error("error in endpoint.GetProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.GetProduct: %w", err)
	}

	return product, nil
}

func (e *Endpoint) AddProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := e.service.AddProduct(ctx, prod)
	if err != nil {
		e.log.Error("error in endpoint.AddProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.AddProduct: %w", err)
	}

	return product, nil
}

func (e *Endpoint) DeleteProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	res, err := e.service.DeleteProduct(ctx, id)
	if err != nil {
		e.log.Error("error in endpoint.DeleteProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.DeleteProduct: %w", err)
	}

	return &product_grpc.ProductResponse{Deleted: res.GetDeleted()}, nil
}

func (e *Endpoint) UpdateProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := e.service.UpdateProduct(ctx, prod)
	if err != nil {
		e.log.Error("error in endpoint.UpdateProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.UpdateProduct: %w", err)
	}

	return product, nil
}
