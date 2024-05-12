package product

import (
	product_grpc "client/api/note_v1"
	"client/internal/metric"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type ServiceInterface interface {
	GetProducts(ctx context.Context) (*product_grpc.AllProductMessage, error)
	GetProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error)
	AddProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
	DeleteProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error)
	UpdateProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
	metric  *metric.Metric
}

func NewEndpoint(service *Service, log *slog.Logger, metric *metric.Metric) *Endpoint {
	return &Endpoint{service: service, log: log, metric: metric}
}

func (e *Endpoint) GetProducts(writer http.ResponseWriter, request *http.Request) {

	start := time.Now()
	defer func() {
		method := request.Method
		elapsed := time.Since(start).Seconds()
		e.metric.RequestCounter.WithLabelValues(method).Inc()
		e.metric.RequestHistogram.WithLabelValues(method).Observe(elapsed)
	}()

	products, err := e.service.GetProducts(request.Context())
	if err != nil {
		e.log.Error("error in endpoint.GetProducts: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(&products); err != nil {
		e.log.Error("error in endpoint.GetProducts: " + err.Error())
	}
}

func (e *Endpoint) GetProduct(writer http.ResponseWriter, request *http.Request) {
	var productID product_grpc.ProductRequest

	if err := json.NewDecoder(request.Body).Decode(&productID); err != nil {
		e.log.Error("error in enpoint.GetProduct: " + err.Error())
	}

	product, err := e.service.GetProduct(request.Context(), &productID)
	if err != nil {
		e.log.Error("error in enpoint.GetProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(&product); err != nil {
		e.log.Error("error in enpoint.GetProduct: " + err.Error())
	}
}

func (e *Endpoint) AddProduct(writer http.ResponseWriter, request *http.Request) {
	var prod product_grpc.ProductMessage

	if err := json.NewDecoder(request.Body).Decode(&prod); err != nil {
		e.log.Error("error in enpoint.AddProduct: " + err.Error())
	}

	product, err := e.service.AddProduct(request.Context(), &prod)
	if err != nil {
		e.log.Error("error in enpoint.AddProduct: " + err.Error())
	}

	if err := json.NewEncoder(writer).Encode(product); err != nil {
		e.log.Error("error in enpoint.AddProduct: " + err.Error())
	}
}

func (e *Endpoint) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	var productID product_grpc.ProductRequest

	if err := json.NewDecoder(request.Body).Decode(&productID); err != nil {
		e.log.Error("error in Delete.AddProduct: " + err.Error())
	}

	result, err := e.service.DeleteProduct(request.Context(), &productID)
	if err != nil {
		e.log.Error("error in Delete.AddProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(result); err != nil {
		e.log.Error("error in Delete.AddProduct: " + err.Error())
	}
}

func (e *Endpoint) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	var prod product_grpc.ProductMessage

	if err := json.NewDecoder(request.Body).Decode(&prod); err != nil {
		e.log.Error("error in Update.AddProduct: " + err.Error())
	}

	product, err := e.service.UpdateProduct(request.Context(), &prod)
	if err != nil {
		e.log.Error("error in Update.AddProduct: " + err.Error())
	}

	if err = json.NewEncoder(writer).Encode(product); err != nil {
		e.log.Error("error in Update.AddProduct: " + err.Error())
	}
}
