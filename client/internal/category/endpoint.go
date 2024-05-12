package category

import (
	product_grpc "client/api/note_v1"
	"client/internal/metric"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ServiceInterface interface {
	GetCategories(ctx context.Context) (*product_grpc.AllCategoryMessage, error)
	AddCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error)
	UpdateCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error)
	DeleteCategory(ctx context.Context, id *product_grpc.CategoryRequest) (*product_grpc.CategoryResponse, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
	metric  *metric.Metric
}

func NewEndpoint(service *Service, log *slog.Logger, metric *metric.Metric) *Endpoint {
	return &Endpoint{service: service, log: log, metric: metric}
}

func (e *Endpoint) GetCategories(writer http.ResponseWriter, request *http.Request) {
	categories, err := e.service.GetCategories(request.Context())
	if err != nil {
		e.log.Error("error in Endpoint's method GetCategories: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&categories); err != nil {
		e.log.Error("error in Endpoint's method GetCategories: " + err.Error())
	}
}

func (e *Endpoint) AddCategory(writer http.ResponseWriter, request *http.Request) {
	var cat product_grpc.CategoryMessage
	if err := json.NewDecoder(request.Body).Decode(&cat); err != nil {
		e.log.Error("error in Endpoint's method AddCategory: " + err.Error())
	}

	category, err := e.service.AddCategory(request.Context(), &cat)
	if err != nil {
		e.log.Error("error in Endpoint's method AddCategory: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&category); err != nil {
		e.log.Error("error in Endpoint's method AddCategory: " + err.Error())
	}
}

func (e *Endpoint) UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	var cat product_grpc.CategoryMessage
	if err := json.NewDecoder(request.Body).Decode(&cat); err != nil {
		e.log.Error("error in Endpoint's method UpdateCategory: " + err.Error())
	}

	category, err := e.service.UpdateCategory(request.Context(), &cat)
	if err != nil {
		e.log.Error("error in Endpoint's method UpdateCategory: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&category); err != nil {
		e.log.Error("error in Endpoint's method UpdateCategory: " + err.Error())
	}
}

func (e *Endpoint) DeleteCategory(writer http.ResponseWriter, request *http.Request) {
	var categoryID product_grpc.CategoryRequest

	if err := json.NewDecoder(request.Body).Decode(&categoryID); err != nil {
		e.log.Error("error in Endpoint's method DeleteCategory: " + err.Error())
	}

	result, err := e.service.DeleteCategory(request.Context(), &categoryID)
	if err != nil {
		e.log.Error("error in Endpoint's method DeleteCategory: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&result); err != nil {
		e.log.Error("error in Endpoint's method DeleteCategory: " + err.Error())
	}
}
